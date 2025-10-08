package providers

import (
	"github.com/jarusk/goscores/internal/models"
)

type Sport int
type League int

const (
	Baseball Sport = iota
	Hockey
)

func (s Sport) String() string {
	switch s {
	case Hockey:
		return "hockey"
	case Baseball:
		return "baseball"
	default:
		return "unknown"
	}
}

const (
	MLB League = iota
	NHL
)

func (l League) String() string {
	switch l {
	case MLB:
		return "mlb"
	case NHL:
		return "nhl"
	default:
		return "unknown"
	}
}

type Provider interface {
	GetName() string
	GetScores(Sport, League, models.GameDate) []models.Match
}
