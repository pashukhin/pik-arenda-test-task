package service

import "time"

type interval struct {
	start, end time.Time
}

func newInterval(start, end time.Time) *interval {
	if start.Unix() == end.Unix() {
		return nil
	}
	if start.Before(end) {
		return &interval{start, end}
	}
	return &interval{end, start}
}

func newIntervalUnix(start, end int64) *interval {
	if start == end {
		return nil
	}
	if start < end {
		return newIntervalUnix(end, start)
	}
	return newInterval(time.Unix(start, 0), time.Unix(end, 0))
}

func (i *interval) isZeroOrNegative() bool {
	return i.start.Unix() >= i.end.Unix()
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func intersectionAndParts(a, b *interval) (int, left, right *interval) {
	if a == nil {
		if b == nil {
			return
		}
		right = newInterval(b.start, b.end)
		return
	}
	if b == nil {
		left = newInterval(a.start, a.end)
		return
	}
	asu, aeu, bsu, beu := a.start.Unix(), a.end.Unix(), b.start.Unix(), b.end.Unix()
	maxStart := maxInt64(asu, bsu)
	minEnd := minInt64(aeu, beu)
	minStart := minInt64(asu, bsu)
	maxEnd := maxInt64(aeu, beu)
	if (aeu > bsu) && (beu > asu) {
		int = &interval{time.Unix(maxStart, 0), time.Unix(minEnd, 0)}
		if minStart != maxStart {
			left = &interval{time.Unix(minStart, 0), time.Unix(maxStart, 0)}
		}
		if minEnd != maxEnd {
			right = &interval{time.Unix(minEnd, 0), time.Unix(maxEnd, 0)}
		}
		return
	}
	if minStart != minEnd {
		left = &interval{time.Unix(minStart, 0), time.Unix(minEnd, 0)}
	}
	if maxStart != maxEnd {
		right = &interval{time.Unix(maxStart, 0), time.Unix(maxEnd, 0)}
	}
	return
}
