package service

import (
	"github.com/pashukhin/pik-arenda-test-task/entity"
	"time"
)

type FreeSlot struct {
	fs *FreeSchedule
}

func NewFreeSlot(fs *FreeSchedule) *FreeSlot {
	return &FreeSlot{fs}
}

func (s *FreeSlot) List(from, to *time.Time) (list []*entity.FreeSlot, err error) {
	return s.fs.getFreeSlots(from, to)
}