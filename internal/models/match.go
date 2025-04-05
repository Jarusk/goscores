package models

import (
	"fmt"
	"time"
)

type GameDate struct {
	t time.Time
}

func NewGameDate(t *time.Time) GameDate {
	if t == nil {
		x := time.Now()
		t = &x
	}
	return GameDate{t: *t}
}
func (g GameDate) String() string {
	return fmt.Sprintf("%d%02d%02d", g.t.Year(), g.t.Month(), g.t.Day())
}

type MatchState uint8

const (
	Undefined MatchState = iota
	Scheduled
	Ongoing
	Finished
)

type Match struct {
	Home         Team
	Away         Team
	HomePoints   int
	AwayPoints   int
	Segment      int
	SegmentUnit  string
	SegmentClock string
	Start        time.Time
	State        MatchState
}

func (m MatchState) String() string {
	switch m {
	case Scheduled:
		return "Scheduled"
	case Ongoing:
		return "Ongoing"
	case Finished:
		return "Finished"
	}

	return "Undefined"
}
