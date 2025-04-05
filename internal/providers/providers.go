package providers

import (
	"github.com/jarusk/goscore/internal/models"
)

type Sport string
type League string

type Provider interface {
	GetName() string
	GetScores(Sport, League, models.GameDate) []models.Match
}
