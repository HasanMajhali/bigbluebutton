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

class SetPresentationDownloadable extends AbstractPresentationRecordEvent {
  import SetPresentationDownloadable._

  setEvent("SetPresentationDownloadable")

  def setPresentationName(name: String): Unit = {
    eventMap.put(PRES_NAME, name)
  }

  def setDownloadable(value: Boolean): Unit = {
    eventMap.put(DOWNLOADABLE, value.toString)
  }

}

object SetPresentationDownloadable {
  protected final val PRES_NAME = "presentationName"
  protected final val DOWNLOADABLE = "downloadable"
}