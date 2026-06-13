package mqtt

import "fmt"

const (
	TrainEventsTopic = "trains/+/schedules/+/events"
	TrainStopsTopic  = "trains/+/schedules/+/stops"
	TrainDelayTopic  = "trains/+/schedules/+/delay"
)

func TrainEventsTopicFor(trainUUID string, scheduleID int64) string {
	return fmt.Sprintf("trains/%s/schedules/%d/events", trainUUID, scheduleID)
}

func TrainStopTopicFor(trainUUID string, scheduleID int64) string {
	return fmt.Sprintf("trains/%s/schedules/%d/stops", trainUUID, scheduleID)
}

func TrainDelayTopicFor(trainUUID string, scheduleID int64) string {
	return fmt.Sprintf("trains/%s/schedules/%d/delay", trainUUID, scheduleID)
}
