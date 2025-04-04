package org.bigbluebutton.core.db

import org.bigbluebutton.core.apps.ScreenshareModel
import org.bigbluebutton.core.apps.ScreenshareModel._
import org.bigbluebutton.core.util.RandomStringGenerator
import slick.jdbc.PostgresProfile.api._
import slick.lifted.ProvenShape

case class ScreenshareDbModel(
    screenshareId:   String,
    meetingId:       String,
    voiceConf:       String,
    screenshareConf: String,
    contentType:     String,
    stream:          String,
    vidWidth:        Int,
    vidHeight:       Int,
    hasAudio:        Boolean,
    startedAt:       java.sql.Timestamp         = new java.sql.Timestamp(System.currentTimeMillis()),
    stoppedAt:       Option[java.sql.Timestamp]
)

class ScreenshareDbTableDef(tag: Tag) extends Table[ScreenshareDbModel](tag, "screenshare") {
  val screenshareId = column[String]("screenshareId", O.PrimaryKey)
  val meetingId = column[String]("meetingId")
  val voiceConf = column[String]("voiceConf")
  val screenshareConf = column[String]("screenshareConf")
  val contentType = column[String]("contentType")
  val stream = column[String]("stream")
  val vidWidth = column[Int]("vidWidth")
  val vidHeight = column[Int]("vidHeight")
  val hasAudio = column[Boolean]("hasAudio")
  val startedAt = column[java.sql.Timestamp]("startedAt")
  val stoppedAt = column[Option[java.sql.Timestamp]]("stoppedAt")
  override def * : ProvenShape[ScreenshareDbModel] = (
    screenshareId, meetingId, voiceConf, screenshareConf, contentType, stream, vidWidth, vidHeight, hasAudio, startedAt, stoppedAt
  ).<>(ScreenshareDbModel.tupled, ScreenshareDbModel.unapply)
}

object ScreenshareDAO {
  def insert(meetingId: String, screenshareModel: ScreenshareModel) = {
    DatabaseConnection.enqueue(
      TableQuery[ScreenshareDbTableDef].forceInsert(
        ScreenshareDbModel(
          screenshareId = s"${System.currentTimeMillis()}-${RandomStringGenerator.randomAlphanumericString(8)}",
          meetingId = meetingId,
          voiceConf = getVoiceConf(screenshareModel),
          screenshareConf = getScreenshareConf(screenshareModel),
          contentType = getContentType(screenshareModel),
          stream = getRTMPBroadcastingUrl(screenshareModel),
          vidWidth = getScreenshareVideoWidth(screenshareModel),
          vidHeight = getScreenshareVideoHeight(screenshareModel),
          hasAudio = getHasAudio(screenshareModel),
          startedAt = new java.sql.Timestamp(System.currentTimeMillis()),
          stoppedAt = None
        )
      )
    )
  }

  def updateStopped(meetingId: String, stream: String) = {
    DatabaseConnection.enqueue(
      TableQuery[ScreenshareDbTableDef]
        .filter(_.meetingId === meetingId)
        .filter(_.stream === stream)
        .filter(_.stoppedAt.isEmpty)
        .map(ev => ev.stoppedAt)
        .update(Some(new java.sql.Timestamp(System.currentTimeMillis())))
    )
  }

}