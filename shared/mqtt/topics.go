package mqtt

import "fmt"

const (
	TrainEventsTopic = "trains/+/events"
	TrainStopsTopic  = "trains/+/stops/+"
	TrainDelayTopic  = "trains/+/delay"
)

func TrainEventsTopicFor(trainUUID string) string {
	return fmt.Sprintf("trains/%s/events", trainUUID)
}

func TrainStopTopicFor(trainUUID string, stopID int64) string {
	return fmt.Sprintf("trains/%s/stops/%d", trainUUID, stopID)
}
