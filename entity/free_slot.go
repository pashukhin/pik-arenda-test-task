package entity

import "time"

type FreeSlot struct {
	Start time.Time `json:"start"`
	End time.Time `json:"end"`
}
