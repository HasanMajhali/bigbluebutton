/**
 * BigBlueButton open source conferencing system - http://www.bigbluebutton.org/
 *
 * Copyright (c) 2017 BigBlueButton Inc. and by respective authors (see below).
 *
 * This program is free software; you can redistribute it and/or modify it under the
 * terms of the GNU Lesser General Public License as published by the Free Software
 * Foundation; either version 3.0 of the License, or (at your option) any later
 * version.
 *
 * BigBlueButton is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
 * PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License along
 * with BigBlueButton; if not, see <http://www.gnu.org/licenses/>.
 *
 */

package org.bigbluebutton.core.record.events

class ActivateTimerRecordEvent extends AbstractTimerRecordEvent {
  import ActivateTimerRecordEvent._

  setEvent("ActivateTimerEvent")

  def setStopwatch(value: Boolean): Unit = {
    eventMap.put(STOPWATCH, value.toString)
  }

  def setRunning(value: Boolean): Unit = {
    eventMap.put(RUNNING, value.toString)
  }

  def setTime(value: Int): Unit = {
    eventMap.put(TIME, value.toString)
  }

  def setAccumulated(value: Int): Unit = {
    eventMap.put(ACCUMULATED, value.toString)
  }

  def setTimestamp(value: Int): Unit = {
    eventMap.put(TIMESTAMP, value.toString)
  }

  def setTrack(value: String): Unit = {
    eventMap.put(TRACK, value)
  }
}

object ActivateTimerRecordEvent {
  protected final val STOPWATCH = "stopwatch"
  protected final val RUNNING = "running"
  protected final val TIME = "time"
  protected final val ACCUMULATED = "accumulated"
  protected final val TIMESTAMP = "timestamp"
  protected final val TRACK = "track"
}
