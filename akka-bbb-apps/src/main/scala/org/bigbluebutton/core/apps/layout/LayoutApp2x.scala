package org.bigbluebutton.core.apps.layout

import org.bigbluebutton.core.running.MeetingActor

trait LayoutApp2x
  extends BroadcastLayoutMsgHdlr
  with BroadcastPushLayoutMsgHdlr
  with GetCurrentLayoutReqMsgHdlr {

  this: MeetingActor =>
}
