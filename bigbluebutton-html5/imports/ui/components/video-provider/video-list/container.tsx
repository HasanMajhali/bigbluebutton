import React, {
  useContext,
  useEffect,
  useState,
} from 'react';
import VideoList from '/imports/ui/components/video-provider/video-list/component';
import { layoutSelect, layoutDispatch } from '/imports/ui/components/layout/context';
import { useNumberOfPages } from '/imports/ui/components/video-provider/hooks';
import { VideoItem } from '/imports/ui/components/video-provider/types';
import { Layout, Output } from '/imports/ui/components/layout/layoutTypes';
import { PluginsContext } from '/imports/ui/components/components-data/plugin-context/context';
import { UpdatedDataForUserCameraDomElement } from 'bigbluebutton-html-plugin-sdk/dist/cjs/dom-element-manipulation/user-camera/types';
import { HookEvents } from 'bigbluebutton-html-plugin-sdk/dist/cjs/core/enum';
import { DomElementManipulationHooks } from 'bigbluebutton-html-plugin-sdk/dist/cjs/dom-element-manipulation/enums';
import { UpdatedEventDetails } from 'bigbluebutton-html-plugin-sdk/dist/cjs/core/types';

interface VideoListContainerProps {
  streams: VideoItem[];
  currentVideoPageIndex: number;
  cameraDock: Output['cameraDock'];
  focusedId: string;
  handleVideoFocus: (id: string) => void;
  isGridEnabled: boolean;
  onVideoItemMount: (stream: string, video: HTMLVideoElement) => void;
  onVideoItemUnmount: (stream: string) => void;
  onVirtualBgDrop: (stream: string, type: string, name: string, data: string) => Promise<unknown>;
}

const VideoListContainer: React.FC<VideoListContainerProps> = (props) => {
  const layoutType = layoutSelect((i: Layout) => i.layoutType);
  const layoutContextDispatch = layoutDispatch();
  const {
    streams,
    cameraDock,
    currentVideoPageIndex,
    focusedId,
    handleVideoFocus,
    isGridEnabled,
    onVideoItemMount,
    onVideoItemUnmount,
    onVirtualBgDrop,
  } = props;
  const numberOfPages = useNumberOfPages();

  const { domElementManipulationIdentifiers } = useContext(PluginsContext);

  const [userCamerasRequestedFromPlugin, setUserCamerasRequestedFromPlugin] = useState<
    UpdatedDataForUserCameraDomElement[]>([]);
  useEffect(() => {
    const dataToSend = userCamerasRequestedFromPlugin.filter((
      userCamera,
    ) => domElementManipulationIdentifiers.USER_CAMERA?.includes(userCamera.streamId));
    window.dispatchEvent(
      new CustomEvent<UpdatedEventDetails<UpdatedDataForUserCameraDomElement[]>>(HookEvents.BBB_CORE_SENT_NEW_DATA, {
        detail: {
          hook: DomElementManipulationHooks.USER_CAMERA,
          data: dataToSend,
        },
      }),
    );
  }, [domElementManipulationIdentifiers, userCamerasRequestedFromPlugin]);
  return (
    !streams.length
      ? null
      : (
        <VideoList
          layoutType={layoutType}
          setUserCamerasRequestedFromPlugin={setUserCamerasRequestedFromPlugin}
          layoutContextDispatch={layoutContextDispatch}
          numberOfPages={numberOfPages}
          currentVideoPageIndex={currentVideoPageIndex}
          cameraDock={cameraDock}
          focusedId={focusedId}
          handleVideoFocus={handleVideoFocus}
          isGridEnabled={isGridEnabled}
          streams={streams}
          onVideoItemMount={onVideoItemMount}
          onVideoItemUnmount={onVideoItemUnmount}
          onVirtualBgDrop={onVirtualBgDrop}
        />
      )
  );
};

export default VideoListContainer;
