package org.bigbluebutton.core2.message.handlers.guests

import org.bigbluebutton.common2.msgs.SetPrivateGuestLobbyMessageCmdMsg
import org.bigbluebutton.core.apps.{ PermissionCheck, RightsManagementTrait }
import org.bigbluebutton.core.db.UserStateDAO
import org.bigbluebutton.core.models.GuestsWaiting
import org.bigbluebutton.core.running.{ LiveMeeting, MeetingActor, OutMsgRouter }
import org.bigbluebutton.core.util.HtmlUtil.htmlToHtmlEntities
import org.bigbluebutton.core2.message.senders.MsgBuilder

trait SetPrivateGuestLobbyMessageCmdMsgHdlr extends RightsManagementTrait {
  this: MeetingActor =>

  val liveMeeting: LiveMeeting
  val outGW: OutMsgRouter

  def handleSetPrivateGuestLobbyMessageCmdMsg(msg: SetPrivateGuestLobbyMessageCmdMsg): Unit = {
    if (permissionFailed(PermissionCheck.MOD_LEVEL, PermissionCheck.VIEWER_LEVEL, liveMeeting.users2x, msg.header.userId)) {
      val meetingId = liveMeeting.props.meetingProp.intId
      val reason = "No permission to send private guest lobby messages."
      PermissionCheck.ejectUserForFailedPermission(meetingId, msg.header.userId, reason, outGW, liveMeeting)
    } else {
      val sanitizedMessage = htmlToHtmlEntities(msg.body.message)
      GuestsWaiting.setPrivateGuestLobbyMessage(liveMeeting.guestsWaiting, msg.body.guestId, sanitizedMessage)
      UserStateDAO.updateGuestLobbyMessage(msg.header.meetingId, msg.body.guestId, sanitizedMessage)
      val event = MsgBuilder.buildPrivateGuestLobbyMsgChangedEvtMsg(
        liveMeeting.props.meetingProp.intId,
        msg.header.userId,
        msg.body.guestId,
        sanitizedMessage
      )
      outGW.send(event)
    }
  }
}
