package providers

import (
	"time"

	"github.com/jarusk/goscore/internal/models"
)

type Sport string
type League string

type Provider interface {
	GetName() string
	GetScores(Sport, League, time.Time) []models.Match
}
