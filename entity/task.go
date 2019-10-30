package entity

import "time"

type Task struct {
	ID       int   `json:"id" db:"id"`
	WorkerID     *int  `json:"worker_id" db:"worker_id"`
	Start time.Time `json:"start" db:"start"`
	End time.Time `json:"end" db:"end"`
	Cancelled bool `json:"cancelled" db:"cancelled"`
}
