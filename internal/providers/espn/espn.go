package espn

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jarusk/goscore/internal/models"
	"github.com/jarusk/goscore/internal/providers"
)

// Compile-time interface check
var _ providers.Provider = &EspnProvider{}

type EspnProvider struct {
	Client *http.Client
}

// GetName implements providers.Provider.
func (e *EspnProvider) GetName() string {
	return "ESPN"
}

// GetScores implements providers.Provider.
func (e *EspnProvider) GetScores(sport providers.Sport, league providers.League, date time.Time) []models.Match {

	matches := []models.Match{}

	scoreboardURL := getScoreboardURL(sport, league)
	resp, err := e.Client.Get(scoreboardURL)
	if err != nil {
		slog.Warn("failed to get scores",
			"provider", e.GetName(),
			"sport", sport,
			"league", league,
			"error", err,
		)
		return matches
	}
	defer resp.Body.Close()

	var parsed scoreboard

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("failed to read body to byte array",
			"provider", e.GetName(),
			"sport", sport,
			"league", league,
			"error", err,
		)
	}
	err = json.Unmarshal(body, &parsed)

	if err != nil {
		slog.Error("failed to parse response",
			"provider", e.GetName(),
			"sport", sport,
			"league", league,
			"error", err,
		)

		return matches
	}

	for _, event := range parsed.Events {
		m := models.Match{}
		m.Start = parseTimestamp(event.Date)

		for _, competition := range event.Competitions {
			for _, competitor := range competition.Competitors {
				parsedScore, err := strconv.Atoi(competitor.Score)
				if err != nil {
					slog.Warn("failed to parse score",
						"raw", competitor.Score,
						"err", err,
					)
				}
				parsedTeam := models.Team{Name: competitor.Team.DisplayName, Slug: competitor.Team.Abbreviation}

				switch competitor.HomeAway {
				case "home":
					m.Home = parsedTeam
					m.HomePoints = parsedScore
				case "away":
					m.Away = parsedTeam
					m.AwayPoints = parsedScore
				}
			}

			m.Segment = competition.Status.Period
			m.SegmentUnit = "period"

			parsedState, err := strconv.Atoi(competition.Status.Type.ID)
			if err != nil {
				slog.Warn("failed to parse match state",
					"err", err,
				)
			}

			m.State = models.MatchState(parsedState)
		}

		matches = append(matches, m)
	}

	return matches
}

func parseTimestamp(timestamp string) time.Time {
	patched := strings.ReplaceAll(timestamp, "Z", ":00Z")

	start, err := time.Parse(time.RFC3339, patched)
	if err != nil {
		slog.Warn("failed to parse start time",
			"raw", timestamp,
			"patched", patched,
			"err", err,
		)
	}
	return start
}
