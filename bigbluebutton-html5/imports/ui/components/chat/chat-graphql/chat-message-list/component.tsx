import React, {
  useCallback,
  useEffect,
  useState,
  useMemo,
  KeyboardEventHandler,
} from 'react';
import { makeVar, useMutation } from '@apollo/client';
import { defineMessages, useIntl } from 'react-intl';
import useChat from '/imports/ui/core/hooks/useChat';
import useIntersectionObserver from '/imports/ui/hooks/useIntersectionObserver';
import { ChatEvents } from '/imports/ui/core/enums/chat';
import { layoutSelect } from '/imports/ui/components/layout/context';
import { Layout } from '/imports/ui/components/layout/layoutTypes';
import { Message } from '/imports/ui/Types/message';
import ChatListPage from './page/component';
import LAST_SEEN_MUTATION from './queries';
import {
  MessageList,
  UnreadButton,
  PageWrapper,
} from './styles';
import useReactiveRef from '/imports/ui/hooks/useReactiveRef';
import useStickyScroll from '/imports/ui/hooks/useStickyScroll';
import ChatReplyIntention from '../chat-reply-intention/component';
import ChatEditingWarning from '../chat-editing-warning/component';
import KEY_CODES from '/imports/utils/keyCodes';
import useMeeting from '/imports/ui/core/hooks/useMeeting';
import useCurrentUser from '/imports/ui/core/hooks/useCurrentUser';
import {
  useIsReplyChatMessageEnabled,
  useIsChatMessageReactionsEnabled,
  useIsEditChatMessageEnabled,
  useIsDeleteChatMessageEnabled,
} from '/imports/ui/services/features';
import { CHAT_DELETE_REACTION_MUTATION, CHAT_SEND_REACTION_MUTATION } from './page/chat-message/mutations';
import logger from '/imports/startup/client/logger';
import { ChatLoading } from '../component';

const PAGE_SIZE = 50;
const CLEANUP_TIMEOUT = 3000;

const intlMessages = defineMessages({
  loadMoreButtonLabel: {
    id: 'app.chat.loadMoreButtonLabel',
    description: 'Label for load more button',
  },
  moreMessages: {
    id: 'app.chat.moreMessages',
    description: 'Chat message when the user has unread messages below the scroll',
  },
});

interface ChatListProps {
  totalUnread: number;
  totalPages: number;
  chatId: string;
  isRTL: boolean;
  setMessageAsSeenMutation: (
    data: {
      variables: {
        chatId: string,
        lastSeenAt: string,
      },
    }
  ) => void;
}

const isElement = (el: unknown): el is HTMLElement => {
  return el instanceof HTMLElement;
};

const isMap = (map: unknown): map is Map<number, string> => {
  return map instanceof Map;
};
let elHeight = 0;

const scrollObserver = new ResizeObserver((entries) => {
  entries.forEach((entry) => {
    const el = entry.target;
    if (isElement(el) && isElement(el.parentElement)) {
      if (el.offsetHeight > elHeight) {
        elHeight = el.offsetHeight;
        el.parentElement.scrollTop = el.parentElement.scrollHeight + el.parentElement.clientHeight;
      } else {
        elHeight = 0;
      }
    }
  });
});

const setLastSender = (lastSenderPerPage: Map<number, string>) => {
  return (page: number, sender: string) => {
    if (isMap(lastSenderPerPage)) {
      lastSenderPerPage.set(page, sender);
    }
  };
};

const lastSeenQueue = makeVar<{ [key: string]: Set<number> }>({});
const setter = makeVar<{ [key: string]:(lastSeenTime: string) => void }>({});
const lastSeenAtVar = makeVar<{ [key: string]: number }>({});
const chatIdVar = makeVar<string>('');

const dispatchLastSeen = () => setTimeout(() => {
  const lastSeenQueueValue = lastSeenQueue();
  if (lastSeenQueueValue[chatIdVar()]) {
    const lastTimeQueue = Array.from(lastSeenQueueValue[chatIdVar()]);
    const lastSeenTime = Math.max(...lastTimeQueue);
    const lastSeenAtVarValue = lastSeenAtVar();
    if (lastSeenTime > (lastSeenAtVarValue[chatIdVar()] ?? 0)) {
      lastSeenAtVar({ ...lastSeenAtVar(), [chatIdVar()]: lastSeenTime });
      setter()[chatIdVar()](new Date(lastSeenTime).toISOString());
    }
  }
}, 500);

const findFirstFocusableChild = (element: HTMLElement) => {
  const childElements = Array.from(element.children) as HTMLElement[];
  return childElements.find((c) => (c as HTMLElement).dataset.focusable === 'true');
};

const findLastFocusableChild = (element: HTMLElement) => {
  const childElements = Array.from(element.children) as HTMLElement[];
  return childElements.reverse().find((c) => (c as HTMLElement).dataset.focusable === 'true');
};

const roving = (
  event: React.KeyboardEvent<HTMLElement>,
  elementsList: HTMLElement,
  setFocused?: (el?: HTMLElement | null) => void,
) => {
  const numberOfChilds = elementsList.childElementCount;
  const isTargetAChild = event.target instanceof HTMLDivElement && event.target.dataset.focusable === 'true';
  const isTargetElement = event.target === event.currentTarget;
  const { activeElement } = document;
  const element = activeElement?.classList.contains('chat-message-container')
    ? activeElement
    : null;

  if (!(isTargetAChild || isTargetElement)) return;

  if (event.keyCode === KEY_CODES.ESCAPE && isTargetAChild) {
    elementsList.focus();
  }

  if (event.keyCode === KEY_CODES.ARROW_DOWN) {
    event.preventDefault();
    const firstElement = findFirstFocusableChild(elementsList);
    let elRef = element && numberOfChilds > 1 ? (element.nextSibling as HTMLElement) : firstElement;

    while (elRef && elRef.dataset.focusable !== 'true' && elRef.nextSibling) {
      elRef = elRef.nextSibling as HTMLElement;
    }

    elRef = (elRef && elRef.dataset.focusable === 'true') ? elRef : firstElement;
    elRef?.focus?.();
    setFocused?.(elRef);
  }

  if (event.keyCode === KEY_CODES.ARROW_UP) {
    event.preventDefault();
    const lastElement = findLastFocusableChild(elementsList);
    let elRef = element ? (element.previousSibling as HTMLElement) : lastElement;

    while (elRef && elRef.dataset.focusable !== 'true' && elRef.previousSibling) {
      elRef = elRef.previousSibling as HTMLElement;
    }

    elRef = (elRef && elRef.dataset.focusable === 'true') ? elRef : lastElement;
    elRef?.focus?.();
    setFocused?.(elRef);
  }

  if ([KEY_CODES.SPACE, KEY_CODES.ENTER].includes(event.keyCode) && isTargetElement) {
    const elRef = document.activeElement?.firstChild as HTMLElement;
    elRef?.focus?.();
    setFocused?.(elRef);
  }
};

const ChatMessageList: React.FC<ChatListProps> = ({
  totalPages,
  chatId,
  setMessageAsSeenMutation,
  totalUnread,
  isRTL,
}) => {
  const intl = useIntl();
  // I used a ref here because I don't want to re-render the component when the last sender changes
  const lastSenderPerPage = React.useRef<Map<number, string>>(new Map());
  const endSentinelRef = React.useRef<HTMLDivElement | null>(null);
  const startSentinelRef = React.useRef<HTMLDivElement | null>(null);
  const {
    ref: messageListContainerRef,
    current: currentMessageListContainer,
  } = useReactiveRef<HTMLDivElement>(null);
  const messageListRef = React.useRef<HTMLDivElement>(null);
  const cleanupTimeoutRef = React.useRef<ReturnType<typeof setTimeout> | null>(null);
  const [userLoadedBackUntilPage, setUserLoadedBackUntilPage] = useState<number | null>(null);
  const [lastMessageCreatedAt, setLastMessageCreatedAt] = useState<string>('');
  const [followingTail, setFollowingTail] = React.useState(true);
  const [showStartSentinel, setShowStartSentinel] = React.useState(false);
  const [focusedMessageElement, setFocusedMessageElement] = React.useState<HTMLElement | null>();
  const [loadingPages, setLoadingPages] = React.useState(new Set<number>());
  const allPagesLoaded = loadingPages.size === 0;
  const {
    childRefProxy: endSentinelRefProxy,
    intersecting: isEndSentinelVisible,
    parentRefProxy: endSentinelParentRefProxy,
  } = useIntersectionObserver(messageListContainerRef, endSentinelRef);
  const {
    childRefProxy: startSentinelRefProxy,
    intersecting: isStartSentinelVisible,
    parentRefProxy: startSentinelParentRefProxy,
  } = useIntersectionObserver(messageListContainerRef, startSentinelRef);
  const {
    startObserving: startObservingStickyScroll,
    stopObserving: stopObservingStickyScroll,
  } = useStickyScroll(currentMessageListContainer, currentMessageListContainer, 'ne');
  const { data: meeting } = useMeeting((m) => ({
    lockSettings: m?.lockSettings,
    isBreakout: m?.isBreakout,
  }));
  const { data: currentUser } = useCurrentUser((c) => ({
    isModerator: c?.isModerator,
    userLockSettings: c?.userLockSettings,
    locked: c?.locked,
    userId: c?.userId,
  }));
  const CHAT_REPLY_ENABLED = useIsReplyChatMessageEnabled();
  const CHAT_REACTIONS_ENABLED = useIsChatMessageReactionsEnabled();
  const CHAT_EDIT_ENABLED = useIsEditChatMessageEnabled();
  const CHAT_DELETE_ENABLED = useIsDeleteChatMessageEnabled();
  const messageToolbarIsEnabled = [
    CHAT_REPLY_ENABLED,
    CHAT_REACTIONS_ENABLED,
    CHAT_EDIT_ENABLED,
    CHAT_DELETE_ENABLED,
  ].some((config) => config);

  const [chatSendReaction] = useMutation(CHAT_SEND_REACTION_MUTATION);
  const [chatDeleteReaction] = useMutation(CHAT_DELETE_REACTION_MUTATION);

  const sendReaction = useCallback((
    reactionEmoji: string,
    reactionEmojiId: string,
    chatId: string,
    messageId: string,
  ) => {
    chatSendReaction({
      variables: {
        chatId,
        messageId,
        reactionEmoji,
        reactionEmojiId,
      },
    }).catch((e) => {
      logger.error({
        logCode: 'chat_send_message_reaction_error',
        extraInfo: {
          errorName: e?.name,
          errorMessage: e?.message,
        },
      }, `Sending reaction failed: ${e?.message}`);
    });
  }, [chatSendReaction]);

  const deleteReaction = useCallback((
    reactionEmoji: string,
    reactionEmojiId: string,
    chatId: string,
    messageId: string,
  ) => {
    chatDeleteReaction({
      variables: {
        chatId,
        messageId,
        reactionEmoji,
        reactionEmojiId,
      },
    }).catch((e) => {
      logger.error({
        logCode: 'chat_delete_message_reaction_error',
        extraInfo: {
          errorName: e?.name,
          errorMessage: e?.message,
        },
      }, `Deleting reaction failed: ${e?.message}`);
    });
  }, [chatDeleteReaction]);

  useEffect(() => {
    if (isEndSentinelVisible) {
      startObservingStickyScroll();
    } else {
      stopObservingStickyScroll();
    }
    toggleFollowingTail(isEndSentinelVisible);
  }, [isEndSentinelVisible]);

  useEffect(() => {
    if (isStartSentinelVisible) {
      if (followingTail) {
        toggleFollowingTail(false);
      }
      if (!allPagesLoaded) return;
      setUserLoadedBackUntilPage((prev) => {
        if (typeof prev === 'number' && prev > 0) {
          return prev - 1;
        }
        return prev;
      });
    }
  }, [isStartSentinelVisible]);

  useEffect(() => {
    setter({
      ...setter(),
      [chatId]: setLastMessageCreatedAt,
    });
    chatIdVar(chatId);
    setLastMessageCreatedAt('');
  }, [chatId]);

  useEffect(() => {
    if (lastMessageCreatedAt !== '') {
      setMessageAsSeenMutation({
        variables: {
          chatId,
          lastSeenAt: lastMessageCreatedAt,
        },
      });
    }
  }, [lastMessageCreatedAt]);

  useEffect(() => {
    const handler = (e: Event) => {
      if (e instanceof CustomEvent) {
        toggleFollowingTail(false);
        setUserLoadedBackUntilPage(Math.ceil(e.detail.sequence / PAGE_SIZE) - 1);
      }
    };

    window.addEventListener(ChatEvents.CHAT_FOCUS_MESSAGE_REQUEST, handler);

    return () => {
      window.removeEventListener(ChatEvents.CHAT_FOCUS_MESSAGE_REQUEST, handler);
    };
  }, []);

  const markMessageAsSeen = useCallback((message: Message) => {
    if (new Date(message.createdAt).getTime() > new Date((lastMessageCreatedAt || 0)).getTime()) {
      dispatchLastSeen();
      const lastSeenQueueValue = lastSeenQueue();
      if (lastSeenQueueValue[chatId]) {
        lastSeenQueueValue[chatId].add(new Date(message.createdAt).getTime());
        lastSeenQueue(lastSeenQueueValue);
      } else {
        lastSeenQueueValue[chatId] = new Set([new Date(message.createdAt).getTime()]);
        lastSeenQueue(lastSeenQueueValue);
      }
    }
  }, [lastMessageCreatedAt, chatId]);

  const setScrollToTailEventHandler = () => {
    toggleFollowingTail(isEndSentinelVisible);
  };

  const toggleFollowingTail = (toggle: boolean) => {
    setFollowingTail(toggle);
    if (isElement(messageListRef.current)) {
      if (toggle) {
        scrollObserver.observe(messageListRef.current);
      } else {
        if (userLoadedBackUntilPage === null) {
          setUserLoadedBackUntilPage(Math.max(totalPages - 2, 0));
        }
        scrollObserver.unobserve(messageListRef.current);
      }
    }
  };

  const renderUnreadNotification = useMemo(() => {
    if (totalUnread && !followingTail) {
      return (
        <UnreadButton
          aria-hidden="true"
          color="primary"
          size="sm"
          key="unread-messages"
          label={intl.formatMessage(intlMessages.moreMessages)}
          onClick={() => {
            if (endSentinelRef.current) {
              endSentinelRef.current.scrollIntoView({ behavior: 'smooth' });
            }
          }}
        />
      );
    }
    return null;
  }, [totalUnread, followingTail]);

  useEffect(() => {
    const scrollToTailEventHandler = () => {
      if (isElement(endSentinelRef.current)) {
        endSentinelRef.current.scrollIntoView();
      }
    };

    window.addEventListener(ChatEvents.SENT_MESSAGE, scrollToTailEventHandler);

    return () => {
      window.removeEventListener(ChatEvents.SENT_MESSAGE, scrollToTailEventHandler);
    };
  }, []);

  useEffect(() => {
    if (followingTail) {
      if (userLoadedBackUntilPage !== null) {
        cleanupTimeoutRef.current = setTimeout(() => {
          setUserLoadedBackUntilPage(null);
        }, CLEANUP_TIMEOUT);
      }
    } else if (cleanupTimeoutRef.current !== null) {
      clearTimeout(cleanupTimeoutRef.current);
      cleanupTimeoutRef.current = null;
    }
  }, [followingTail, userLoadedBackUntilPage]);

  useEffect(() => {
    if (isElement(endSentinelRef.current)) {
      endSentinelRef.current.scrollIntoView();
    }
  }, []);

  const firstPageToLoad = userLoadedBackUntilPage !== null
    ? userLoadedBackUntilPage
    : Math.max(totalPages - 2, 0);
  const pagesToLoad = (totalPages - firstPageToLoad) || 1;

  const rove: KeyboardEventHandler<HTMLElement> = (e) => {
    if (messageListRef.current) {
      roving(
        e,
        messageListRef.current,
        setFocusedMessageElement,
      );
    }
  };

  const hasMessageToolbar = CHAT_DELETE_ENABLED
    || CHAT_EDIT_ENABLED
    || CHAT_REPLY_ENABLED
    || CHAT_REACTIONS_ENABLED;

  const updateRefs = useCallback((el: HTMLDivElement | null) => {
    messageListContainerRef.current = el;
    startSentinelParentRefProxy.current = el;
    endSentinelParentRefProxy.current = el;
  }, []);

  const setPageLoading = useCallback((page: number) => {
    setLoadingPages((prev) => {
      prev.add(page);
      return new Set(prev);
    });
  }, [setLoadingPages]);

  const clearPageLoading = useCallback((page: number) => {
    setLoadingPages((prev) => {
      prev.delete(page);
      return new Set(prev);
    });
  }, [setLoadingPages]);

  return (
    <>
      {
        [
          <MessageList
            id="chat-list"
            key="message-list-wrapper"
            onMouseUp={() => {
              setScrollToTailEventHandler();
            }}
            onTouchEnd={() => {
              setScrollToTailEventHandler();
            }}
            onScroll={(e) => {
              if (e.target instanceof HTMLDivElement) {
                const userScrolledUp = Math.ceil(e.target.scrollTop + e.target.clientHeight) < e.target.scrollHeight;
                setShowStartSentinel(userScrolledUp);
              }
            }}
            onWheel={(e) => {
              if (e.deltaY < 0 && isStartSentinelVisible && allPagesLoaded) {
                setUserLoadedBackUntilPage((prev) => {
                  if (typeof prev === 'number' && prev > 0) {
                    return prev - 1;
                  }
                  return prev;
                });
              }
            }}
            data-test="chatMessages"
            isRTL={isRTL}
            ref={updateRefs}
            $hasMessageToolbar={hasMessageToolbar}
          >
            {showStartSentinel && (
              <div
                ref={startSentinelRefProxy}
                style={{
                  height: 1,
                  background: 'none',
                }}
                tabIndex={-1}
                aria-hidden
              />
            )}
            <PageWrapper
              role="list"
              aria-live="polite"
              ref={messageListRef}
              tabIndex={hasMessageToolbar ? 0 : -1}
              onKeyDown={rove}
            >
              {Array.from({ length: pagesToLoad }, (_v, k) => k + (firstPageToLoad)).map((page) => {
                return (
                  <ChatListPage
                    firstPageToLoad={firstPageToLoad}
                    key={`page-${page}`}
                    page={page}
                    pageSize={PAGE_SIZE}
                    setLastSender={() => setLastSender(lastSenderPerPage.current)}
                    lastSenderPreviousPage={page ? lastSenderPerPage.current.get(page - 1) : undefined}
                    chatId={chatId}
                    markMessageAsSeen={markMessageAsSeen}
                    scrollRef={messageListContainerRef}
                    meetingDisablePublicChat={!!meeting?.lockSettings?.disablePublicChat}
                    meetingDisablePrivateChat={!!meeting?.lockSettings?.disablePrivateChat}
                    currentUserDisablePublicChat={!!currentUser?.userLockSettings?.disablePublicChat}
                    currentUserId={currentUser?.userId ?? ''}
                    currentUserIsLocked={!!currentUser?.locked}
                    currentUserIsModerator={!!currentUser?.isModerator}
                    isBreakoutRoom={!!meeting?.isBreakout}
                    messageToolbarIsEnabled={messageToolbarIsEnabled}
                    chatDeleteEnabled={CHAT_DELETE_ENABLED}
                    chatEditEnabled={CHAT_EDIT_ENABLED}
                    chatReactionsEnabled={CHAT_REACTIONS_ENABLED}
                    chatReplyEnabled={CHAT_REPLY_ENABLED}
                    deleteReaction={deleteReaction}
                    sendReaction={sendReaction}
                    focusedSequence={Number(focusedMessageElement?.dataset.sequence)}
                    clearPageLoading={clearPageLoading}
                    setPageLoading={setPageLoading}
                  />
                );
              })}
            </PageWrapper>
            <div
              ref={endSentinelRefProxy}
              style={{
                height: 1,
                background: 'none',
              }}
              tabIndex={-1}
              aria-hidden
            />
          </MessageList>,
          renderUnreadNotification,
          <ChatReplyIntention key="chatReplyIntention" />,
          <ChatEditingWarning key="chatEditingWarning" />,
        ]
      }
    </>
  );
};

const ChatMessageListContainer: React.FC = () => {
  const idChatOpen = layoutSelect((i: Layout) => i.idChatOpen);
  const isRTL = layoutSelect((i: Layout) => i.isRTL);

  const CHAT_CONFIG = window.meetingClientSettings.public.chat;
  const PUBLIC_CHAT_KEY = CHAT_CONFIG.public_id;
  const PUBLIC_GROUP_CHAT_KEY = CHAT_CONFIG.public_group_id;

  const isPublicChat = idChatOpen === PUBLIC_CHAT_KEY;
  const chatId = !isPublicChat ? idChatOpen : PUBLIC_GROUP_CHAT_KEY;
  const { data: chatData, loading: currentChatLoading } = useChat((chat) => {
    return {
      chatId: chat.chatId,
      totalMessages: chat.totalMessages,
      totalUnread: chat.totalUnread,
      lastSeenAt: chat.lastSeenAt,
    };
  }, chatId);
  const currentChat = Array.isArray(chatData) ? chatData[0] : chatData;

  const [setMessageAsSeenMutation] = useMutation(LAST_SEEN_MUTATION);

  if (currentChatLoading || !currentChat) return <ChatLoading isRTL={document.dir === 'rtl'} />;

  const { totalMessages = 0 } = currentChat;
  const totalPages = Math.ceil(totalMessages / PAGE_SIZE);
  return (
    <ChatMessageList
      totalPages={totalPages}
      chatId={chatId}
      setMessageAsSeenMutation={setMessageAsSeenMutation}
      totalUnread={currentChat?.totalUnread || 0}
      isRTL={isRTL}
    />
  );
};

export default ChatMessageListContainer;
