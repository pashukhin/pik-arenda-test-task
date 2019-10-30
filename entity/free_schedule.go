package entity

import "time"

type FreeSchedule struct {
	ID       int   `json:"id" db:"id"`
	Start time.Time `json:"start" db:"start"`
	End time.Time `json:"end" db:"end"`
	Value int `json:"value" db:"value"`
}
