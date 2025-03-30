package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jarusk/goscore/internal/models"
	"github.com/jarusk/goscore/internal/providers/espn"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	defdate := time.Now()
	gamedate := fmt.Sprintf("%d%d%d", defdate.Year(), defdate.Month(), defdate.Day())

	slog.SetLogLoggerLevel(slog.LevelDebug)

	slog.Info("starting goscores with defaults",
		"gamedate", gamedate,
		"sport", "hockey",
		"league", "nhl",
	)

	httpClient := http.Client{}
	provider := espn.EspnProvider{Client: &httpClient}

	slog.Debug("fetching scores",
		"provider", provider.GetName(),
	)
	m := provider.GetScores("hockey", "nhl", time.Now())
	slog.Debug("finished fetching scores")

	printTable(m)

	slog.Debug("exiting")
}

func printTable(matches []models.Match) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Home", "Away", "Status", "Period / Start", "Clock", "Score"})
	for _, m := range matches {
		var gameState string
		var gameClock string

		switch m.Segment {
		case 0:
			gameState = m.Start.Local().Format(time.Kitchen)
			gameClock = "-"
		case 3:
			gameState = "-"
			gameClock = "-"
		default:
			gameState = fmt.Sprint(m.Segment)
			gameClock = m.SegmentClock
		}

		score := fmt.Sprintf("%4s %d - %4s %d", m.Home.Slug, m.HomePoints, m.Away.Slug, m.AwayPoints)

		t.AppendRow(table.Row{m.Home.Slug, m.Away.Slug, m.State, gameState, gameClock, score})
	}
	t.Render()
}
