package models

import "time"

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
