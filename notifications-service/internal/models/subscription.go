package models

type Subscription struct {
	ID         int64  `json:"id,omitempty"`
	UserID     int64  `json:"user_id,omitempty"`
	TrainUUID  string `json:"train_uuid,omitempty"`
	ScheduleID int64  `json:"schedule_id,omitempty"`
}
