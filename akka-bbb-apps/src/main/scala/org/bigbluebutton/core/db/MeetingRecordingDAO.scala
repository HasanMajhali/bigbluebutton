package org.bigbluebutton.core.db

import slick.jdbc.PostgresProfile.api._
import slick.lifted.ProvenShape

case class MeetingRecordingDbModel(
    meetingId:               String,
    startedAt:               java.sql.Timestamp,
    startedBy:               String,
    stoppedAt:               Option[java.sql.Timestamp],
    stoppedBy:               Option[String]
)

class MeetingRecordingDbTableDef(tag: Tag) extends Table[MeetingRecordingDbModel](tag, "meeting_recording") {
  val meetingId = column[String]("meetingId", O.PrimaryKey)
  val startedAt = column[java.sql.Timestamp]("startedAt")
  val startedBy = column[String]("startedBy")
  val stoppedAt = column[Option[java.sql.Timestamp]]("stoppedAt")
  val stoppedBy = column[Option[String]]("stoppedBy")

  //  def fk_meetingId: ForeignKeyQuery[MeetingDbTableDef, MeetingDbModel] = foreignKey("fk_meetingId", meetingId, TableQuery[MeetingDbTableDef])(_.meetingId)

  override def * : ProvenShape[MeetingRecordingDbModel] = (meetingId, startedAt, startedBy, stoppedAt, stoppedBy) .<> (MeetingRecordingDbModel.tupled, MeetingRecordingDbModel.unapply)
}

object MeetingRecordingDAO {
  def insertRecording(meetingId: String, startedBy: String) = {

    //Stop any previous recoding for this meeting before starting a new one
    updateStopped(meetingId, startedBy)

    DatabaseConnection.enqueue(
      TableQuery[MeetingRecordingDbTableDef].forceInsert(
        MeetingRecordingDbModel(
          meetingId = meetingId,
          startedAt = new java.sql.Timestamp(System.currentTimeMillis()),
          startedBy = startedBy,
          stoppedAt = None,
          stoppedBy = None,
        )
      )
    )
  }

  def updateStopped(meetingId: String, stoppedBy: String) = {
    DatabaseConnection.enqueue(
      TableQuery[MeetingRecordingDbTableDef]
        .filter(_.meetingId === meetingId)
        .filter(_.stoppedAt.isEmpty)
        .map(r => (r.stoppedAt, r.stoppedBy))
        .update((Some(new java.sql.Timestamp(System.currentTimeMillis())), Some(stoppedBy)))
    )
  }

}