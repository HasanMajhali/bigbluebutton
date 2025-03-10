const { expect, default: test } = require('@playwright/test');
const { MultiUsers } = require('../user/multiusers');
const e = require('../core/elements');
const c = require('./constants');
const { VIDEO_LOADING_WAIT_TIME, ELEMENT_WAIT_LONGER_TIME, ELEMENT_WAIT_EXTRA_LONG_TIME, ELEMENT_WAIT_TIME } = require('../core/constants');
const util = require('./util');
const { sleep } = require('../core/helpers');
const { getSettings } = require('../core/settings');
const { uploadSinglePresentation } = require('../presentation/util');

class CustomParameters extends MultiUsers {
  constructor(browser, context) {
    super(browser, context);
  }

  async showPublicChatOnLogin() {
    await this.modPage.hasElement(e.actions, 'should display the actions button');
    await this.modPage.wasRemoved(e.hidePublicChat, 'should display the hide public chat element when the public chat is open');
  }

  async recordMeeting() {
    await this.modPage.hasElement(e.recordingIndicator, 'should display the recording indicator');
  }

  async showParticipantsOnLogin() {
    await this.modPage.wasRemoved(e.usersList, 'should not display the users list');
  }

  async clientTitle() {
    const pageTitle = await this.modPage.page.title();
    expect(pageTitle, 'should display the changed name of the client title').toContain(`${c.docTitle} - `);
  }

  async displayBrandingArea() {
    await this.modPage.hasElement(e.userListContent, 'should display the user list on the meeting');
    await this.modPage.hasElement(e.brandingAreaLogo, 'should display the logo on the branding area');
  }

  async shortcuts() {
    // Check the initial shortcuts that can be used right after joining the meeting
    await util.checkShortcutsArray(this.modPage, c.initialShortcuts);
    // Join audio
    await this.modPage.waitAndClick(e.joinAudio);
    await this.modPage.joinMicrophone();
    // Open private chat
    await this.modPage.waitAndClick(e.userListItem);
    const lastUserStartPrivateChat = await this.modPage.getLocator(e.startPrivateChat).last();
    await this.modPage.clickOnLocator(lastUserStartPrivateChat);
    await this.modPage.hasElement(e.hidePrivateChat, 'should display the hide private chat element when the user has the private chat open');
    // Check the later shortcuts that can be used after joining audio and opening private chat
    await util.checkShortcutsArray(this.modPage, c.laterShortcuts);
  }

  async customStyle() {
    await this.modPage.hasElement(e.chatButton, 'should display the chat button');
    const resp = await this.modPage.page.evaluate((elem) => {
      return document.querySelectorAll(elem)[0].offsetHeight == 0;
    }, e.presentationTitle);
    expect(resp, 'should display the different style on the meeting').toBeTruthy();
  }

  async autoSwapLayout() {
    await this.modPage.hasElement(e.actions, 'should display the actions button');
    await this.modPage.waitAndClick(e.minimizePresentation);
    const resp = await this.modPage.page.evaluate((elem) => {
      return document.querySelectorAll(elem)[0].offsetHeight !== 0;
    }, e.restorePresentation);
    expect(resp, 'should display the auto swap layout').toBeTruthy();
  }

  async autoJoin() {
    await this.modPage.hasElement(e.chatMessages, 'should display the chat messages', ELEMENT_WAIT_LONGER_TIME);
    await this.modPage.wasRemoved(e.audioModal, 'should not display the audio modal');
  }

  async listenOnlyMode() {
    await this.modPage.hasElement(e.audioSettingsModal, 'should display the audio settings modal when joining', ELEMENT_WAIT_EXTRA_LONG_TIME);
    await this.modPage.waitAndClick(e.joinEchoTestButton);
    await this.modPage.hasElement(e.establishingAudioLabel, 'should display the audio being established');
    await this.modPage.hasElement(e.unmuteMicButton, 'should display the unmute button when user joins audio');
    await this.modPage.waitAndClick(e.unmuteMicButton);
    await this.modPage.hasElement(e.isTalking, 'should display the is talking indicator, after the audio being established and mic unmuted');
    await this.modPage.leaveAudio();
    await this.modPage.waitAndClick(e.joinAudio);
    await this.modPage.hasElement(e.audioSettingsModal, 'should display the audio settings modal after clicked on the join audio button');
  }

  async forceListenOnly() {
    await this.userPage.wasRemoved(e.audioModal, 'should not display the audio modal after joining meeting');
    await this.userPage.hasElement(e.audioDropdownMenu, 'should display the audio dropdown menu');
    await this.userPage.hasElement(e.unmuteMicButton, 'should display the unmute button when user joins audio / listen only');
    await this.userPage.waitAndClick(e.unmuteMicButton);
    await this.userPage.hasElement(e.joinEchoTestButton, 'should display the join echo test modal when user tries to unmute the mic in listen only mode');
    await this.userPage.waitAndClick(e.closeModal);
    await this.userPage.waitAndClick(e.audioDropdownMenu);
    await this.userPage.hasElement(e.leaveAudio, 'should display leave button in the audio dropdown menu');
  }

  async skipCheck() {
    await this.modPage.waitAndClick(e.microphoneButton);
    await this.modPage.hasElement(e.establishingAudioLabel, 'should establish audio');
    await this.modPage.wasRemoved(e.establishingAudioLabel, 'should not display the audio being established label', ELEMENT_WAIT_LONGER_TIME);
    await this.modPage.hasElement(e.unmuteMicButton, 'should display the unmute button when user joins audio');
    await this.modPage.waitAndClick(e.unmuteMicButton);
    await this.modPage.hasElement(e.isTalking, 'should display the is talking element when mic is unmuted');
  }

  async skipCheckOnFirstJoin() {
    await this.modPage.waitAndClick(e.microphoneButton, ELEMENT_WAIT_LONGER_TIME);
    await this.modPage.hasElement(e.establishingAudioLabel, 'should establish audio');
    await this.modPage.hasElement(e.unmuteMicButton, 'should display the unmute button when user joins audio');
    await this.modPage.waitAndClick(e.unmuteMicButton);
    await this.modPage.hasElement(e.isTalking, 'should display the is talking element after unmuting the mic');
    await this.modPage.leaveAudio();
    await this.modPage.waitAndClick(e.joinAudio);
    await this.modPage.waitAndClick(e.microphoneButton);
    await this.modPage.hasElement(e.audioSettingsModal, 'should display the audio settings modal');
  }

  async skipEchoTestIfPreviousDevice() {
    await this.modPage.joinMicrophone();
    await this.modPage.leaveAudio();
    await this.modPage.waitAndClick(e.joinAudio);
    await this.modPage.waitAndClick(e.microphoneButton);
    await this.modPage.waitForSelector(e.establishingAudioLabel);
    await this.modPage.wasRemoved(e.establishingAudioLabel, ELEMENT_WAIT_LONGER_TIME);
    await this.modPage.hasElement(e.isTalking);
  }

  async bannerText() {
    await this.modPage.hasElemnt(e.actions, 'should display the actions button');
    await this.modPage.hasElement(e.notificationBannerBar, 'should display the notification banner bar with a text');
  }

  async bannerColor(colorToRGB) {
    await this.modPage.hasElement(e.notificationBannerBar, 'should display the notifications banner bar');
    const notificationLocator = this.modPage.getLocator(e.notificationBannerBar);
    const notificationBarColor = await notificationLocator.evaluate((elem) => {
      return getComputedStyle(elem).backgroundColor;
    }, e.notificationBannerBar);
    expect(notificationBarColor, 'should display the banner bar with the chosen color').toBe(colorToRGB);
  }

  async hidePresentationOnJoin() {
    await this.modPage.hasElement(e.actions, 'should display the actions button');
    await this.modPage.hasElement(e.restorePresentation, 'should display the restore presentation button for the moderator');
    await this.userPage.hasElement(e.restorePresentation, 'should display the restore presentation button for the attendee');
    await this.userPage.wasRemoved(e.whiteboard, 'should not display the whiteboard for the attendee');
  }

  async forceRestorePresentationOnNewEvents() {
    const { presentationHidden, pollEnabled } = getSettings();
    await this.modPage.waitForSelector(e.whiteboard, ELEMENT_WAIT_LONGER_TIME);
    if (!presentationHidden) {
      await this.userPage.waitForSelector(e.whiteboard);
      await sleep(1000);  // wait for the whiteboard to be fully loaded and stable (zoom)
      await this.userPage.waitAndClick(e.minimizePresentation);
    }
    await this.userPage.wasRemoved(e.whiteboard, 'should remove the whiteboard element for the attendee when minimized');
    await sleep(1000);  // first minimize of presentation takes longer to be fully applied
    // zoom in
    await this.modPage.waitAndClick(e.zoomInButton);
    await this.userPage.hasElement(e.whiteboard, 'should restore presentation when zooming in the slide');
    await this.userPage.hasElement(e.minimizePresentation, 'should display the minimize presentation button when the presentation is restored');
    await this.userPage.waitAndClick(e.minimizePresentation);
    // zoom out
    await this.modPage.waitAndClick(e.zoomOutButton);
    await this.userPage.hasElement(e.whiteboard, 'should restore presentation when zooming out the slide');
    await this.userPage.hasElement(e.minimizePresentation, 'should display the minimize presentation button when the presentation is restored');
    await this.userPage.waitAndClick(e.minimizePresentation);
    // publish polling
    if (pollEnabled) {
      await util.poll(this.modPage, this.userPage);
      await this.userPage.hasElement(e.whiteboard, 'should restore presentation when a poll is posted');
      await this.userPage.hasElement(e.minimizePresentation, 'should display the minimize presentation button when the presentation is restored');
      await this.userPage.waitAndClick(e.minimizePresentation);
    }
    // next slide
    await util.nextSlide(this.modPage);
    await this.userPage.hasElement(e.whiteboard, 'should restore presentation when going to the next slide');
    await this.userPage.hasElement(e.minimizePresentation, 'should display the minimize presentation button when the presentation is restored');
    await this.userPage.waitAndClick(e.minimizePresentation);
    // previous slide
    await util.previousSlide(this.modPage);
    await this.userPage.hasElement(e.whiteboard, 'should restore presentation when going to the previous slide');
    await this.userPage.hasElement(e.minimizePresentation, 'should display the minimize presentation button when the presentation is restored');
    await this.userPage.waitAndClick(e.minimizePresentation);
    // select slide
    await this.modPage.selectSlide('Slide 5');
    await this.userPage.hasElement(e.whiteboard, 'should restore presentation when selecting a different slide');
    await this.userPage.hasElement(e.minimizePresentation, 'should display the minimize presentation button when the presentation is restored');
    await this.userPage.waitAndClick(e.minimizePresentation);
    // new presentation
    await uploadSinglePresentation(this.modPage, e.uploadPresentationFileName, ELEMENT_WAIT_LONGER_TIME);
    await this.userPage.hasElement(e.whiteboard, 'should restore presentation when uploading a new presentation');
    await this.userPage.hasElement(e.minimizePresentation, 'should display the minimize presentation button when the presentation is restored');
    await this.userPage.waitAndClick(e.minimizePresentation);
    // add annotation
    await util.annotation(this.modPage);
    await this.userPage.hasElement(e.whiteboard, 'should restore presentation after the annotation event is triggered');
    await this.userPage.hasElement(e.minimizePresentation, 'should display the minimize presentation button when the presentation is restored');
  }

  async enableVideo() {
    await this.modPage.wasRemoved(e.joinVideo, 'should not display the join video');
  }

  async skipVideoPreview() {
    await this.modPage.shareWebcam(false);
  }

  async skipVideoPreviewOnFirstJoin() {
    await this.modPage.shareWebcam(false);
    await this.modPage.waitAndClick(e.leaveVideo, VIDEO_LOADING_WAIT_TIME);
    await this.modPage.hasElement(e.joinVideo, 'should display the join video button');
    const { videoPreviewTimeout } = this.modPage.settings;
    await this.modPage.shareWebcam(true, videoPreviewTimeout);
  }

  async skipVideoPreviewIfPreviousDevice() {
    await this.modPage.waitForSelector(e.joinVideo);
    const { videoPreviewTimeout } = this.modPage.settings;
    await this.modPage.shareWebcam(true, videoPreviewTimeout);
    await this.modPage.waitAndClick(e.leaveVideo, VIDEO_LOADING_WAIT_TIME);
    await this.modPage.shareWebcam(false);
  }

  async mirrorOwnWebcam() {
    await this.modPage.waitAndClick(e.joinVideo);
    await this.modPage.hasElement(e.webcamMirroredVideoPreview, 'should display the preview of the webcam video being mirroed');
    await this.modPage.waitAndClick(e.startSharingWebcam);
    await this.modPage.hasElement(e.webcamMirroredVideoContainer, 'should display the webcam mirroed video container after the camera is shared');
  }

  async multiUserPenOnly() {
    await this.modPage.waitAndClick(e.multiUsersWhiteboardOn);
    await this.userPage.hasElement(e.wbToolbar);
    const toolsCount = await this.userPage.getSelectorCount(`${e.wbToolbar} button:visible`);
    await expect(toolsCount, 'should display only 1 tool (pen)').toBe(1);
  }

  async presenterTools() {
    await this.modPage.waitForSelector(e.whiteboard, ELEMENT_WAIT_LONGER_TIME);
    await this.modPage.hasElement(e.wbToolbar);
    const toolsCount = await this.modPage.getSelectorCount(`${e.wbToolbar} button:visible`);
    await expect(toolsCount, 'should display only the 3 elements passed in the parameter (select, draw and arrow)').toBe(3);
  }

  async multiUserTools() {
    await this.modPage.waitAndClick(e.multiUsersWhiteboardOn);
    await this.userPage.hasElement(e.wbToolbar);
    const toolsCount = await this.userPage.getSelectorCount(`${e.wbToolbar} button:visible`);
    await expect(toolsCount, 'should display only the 2 elements passed in the parameter (arrow and text)').toBe(2);
  }

  async autoShareWebcam() {
    await this.modPage.hasElement(e.webcamSettingsModal, 'should display the webcam settings modal when auto sharing the webcam ');
  }

  async hideActionsBarTest() {
    await this.modPage.wasRemoved(e.actions, 'should not display the actions button');
    await this.modPage.wasRemoved(e.joinAudio, 'should not display the join audio button');
    await this.modPage.wasRemoved(e.joinVideo, 'should not display the join video button');
    await this.modPage.wasRemoved(e.startScreenSharing, 'should not display the start screensharing');
    await this.modPage.wasRemoved(e.minimizePresentation, 'should not display the minimize presentation button');
    await this.modPage.wasRemoved(e.reactionsButton, 'should not display the reactions button');
  }

  async overrideDefaultLocaleTest() {
    await this.modPage.hasText(e.chatButton, 'Bate-papo público','should display the new overridden default locale');
  }

  async hideNavBarTest() {
    await this.modPage.wasRemoved(e.navbarBackground, 'should not display the navbar');
  }

  async preferredCameraProfileTest() {
    await this.modPage.waitAndClick(e.joinVideo);
    expect(await this.modPage.getLocator(e.selectCameraQualityId).inputValue(), 'should display the selector to choose the camera quality').toBe('low');
    await this.modPage.waitAndClick(e.startSharingWebcam);
  }

  async webcamBackgroundURL() {
    await this.modPage.waitForSelector(e.whiteboard);
    await this.modPage.waitAndClick(e.joinVideo);
    await this.modPage.hasElement(e.webcamSettingsModal, 'should display the webcam settings modal when clicking to join video');
    await this.modPage.waitAndClick(e.backgroundSettingsTitle);
    const appleBackground = await this.modPage.getLocator(e.selectCustomBackground);
    await expect(appleBackground).toHaveCount(1);
    await this.modPage.waitAndClick(e.selectCustomBackground);
    await this.modPage.waitAndClick(e.startSharingWebcam);
    await this.modPage.hasElement(e.webcamMirroredVideoContainer, 'should display the webcam (mirrored) container after successfully sharing webcam');
    const webcamBackgroundURL = this.modPage.getLocator(e.webcamMirroredVideoContainer);
    await expect(webcamBackgroundURL).toHaveScreenshot('webcam-background-passing-url.png');
  }
}

exports.CustomParameters = CustomParameters;
