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

class ResizeAndMoveSlideRecordEvent extends AbstractPresentationRecordEvent {
  import ResizeAndMoveSlideRecordEvent._

  setEvent("ResizeAndMoveSlideEvent")

  def setPresentationName(name: String): Unit = {
    eventMap.put(PRES_NAME, name)
  }

  def setId(id: String): Unit = {
    eventMap.put(ID, id)
  }

  def setXOffset(offset: Double): Unit = {
    eventMap.put(X_OFFSET, offset.toString)
  }

  def setYOffset(offset: Double): Unit = {
    eventMap.put(Y_OFFSET, offset.toString)
  }

  def setWidthRatio(ratio: Double): Unit = {
    eventMap.put(WIDTH_RATIO, ratio.toString)
  }

  def setHeightRatio(ratio: Double): Unit = {
    eventMap.put(HEIGHT_RATIO, ratio.toString)
  }
}

object ResizeAndMoveSlideRecordEvent {
  protected final val PRES_NAME = "presentationName"
  protected final val ID = "id"
  protected final val X_OFFSET = "xOffset"
  protected final val Y_OFFSET = "yOffset"
  protected final val WIDTH_RATIO = "widthRatio"
  protected final val HEIGHT_RATIO = "heightRatio"
}
