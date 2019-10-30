package entity

import "time"

type FreeTimePoint struct {
	ID       int   `json:"id" db:"id"`
	ScheduleID     *int  `json:"schedule_id" db:"schedule_id"`
	TaskID     *int  `json:"task_id" db:"task_id"`
	Point time.Time `json:"point" db:"point"`
	Value int `json:"value" db:"value"`
}
