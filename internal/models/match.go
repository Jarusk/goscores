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
	SegmentClock time.Duration
	Start        time.Time
	State        MatchState
}
