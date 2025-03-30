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
	t.AppendHeader(table.Row{"Home", "Away", "Status", "Start (Local)", "Period", "Clock"})
	for _, m := range matches {
		t.AppendRow(table.Row{m.Home.Slug, m.Away.Slug, m.State, m.Start.Local().Format(time.Kitchen), m.Segment, m.SegmentClock})
	}
	t.Render()
}
