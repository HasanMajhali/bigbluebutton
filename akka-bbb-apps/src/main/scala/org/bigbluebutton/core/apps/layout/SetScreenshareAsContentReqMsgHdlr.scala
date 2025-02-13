package org.bigbluebutton.core.apps.layout

import org.bigbluebutton.ClientSettings.getConfigPropertyValueByPathAsBooleanOrElse
import org.bigbluebutton.common2.msgs._
import org.bigbluebutton.core.apps.{PermissionCheck, RightsManagementTrait}
import org.bigbluebutton.core.db.LayoutDAO
import org.bigbluebutton.core.models.Layouts
import org.bigbluebutton.core.running.OutMsgRouter

trait SetScreenshareAsContentReqMsgHdlr extends RightsManagementTrait {
  this: LayoutApp2x =>

  val outGW: OutMsgRouter

  def handleSetScreenshareAsContentReqMsg(msg: SetScreenshareAsContentReqMsg): Unit = {
    if (liveMeeting.props.meetingProp.disabledFeatures.contains("screenshare")) {
      val meetingId = liveMeeting.props.meetingProp.intId
      val reason = "Screenshare is disabled for this meeting."
      PermissionCheck.ejectUserForFailedPermission(meetingId, msg.header.userId, reason, outGW, liveMeeting)
    } else if (permissionFailed(PermissionCheck.GUEST_LEVEL, PermissionCheck.PRESENTER_LEVEL, liveMeeting.users2x, msg.header.userId)) {
      val meetingId = liveMeeting.props.meetingProp.intId
      val reason = "No permission to set screenshare as content."
      PermissionCheck.ejectUserForFailedPermission(meetingId, msg.header.userId, reason, outGW, liveMeeting)
    } else {
      Layouts.setScreenshareAsContent(liveMeeting.layouts, msg.body.screenshareAsContent)
      LayoutDAO.insertOrUpdate(liveMeeting.props.meetingProp.intId, liveMeeting.layouts)
      sendSetScreenshareAsContentEvtMsg(msg.header.userId)
    }
  }

  def sendSetScreenshareAsContentEvtMsg(fromUserId: String): Unit = {
    val routing = Routing.addMsgToClientRouting(MessageTypes.BROADCAST_TO_MEETING, liveMeeting.props.meetingProp.intId, fromUserId)
    val envelope = BbbCoreEnvelope(SetScreenshareAsContentEvtMsg.NAME, routing)
    val header = BbbClientMsgHeader(SetScreenshareAsContentEvtMsg.NAME, liveMeeting.props.meetingProp.intId, fromUserId)

    val body = SetScreenshareAsContentEvtMsgBody(
      Layouts.getScreenshareAsContent(liveMeeting.layouts),
    )
    val event = SetScreenshareAsContentEvtMsg(header, body)
    val msgEvent = BbbCommonEnvCoreMsg(envelope, event)

    outGW.send(msgEvent)
  }
}
