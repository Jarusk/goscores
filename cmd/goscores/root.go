package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/jarusk/goscores/internal/models"
	"github.com/jarusk/goscores/internal/providers"
	"github.com/jarusk/goscores/internal/providers/espn"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goscores",
	Short: "goscores gives you the latest sports scores",
	Long:  `A tool for fetching the latest sports scores on the CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			fmt.Printf("failed to print help: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(baseballCmd)
	rootCmd.AddCommand(hockeyCmd)
}

func run(sport providers.Sport, league providers.League, date string) {
	g := models.NewGameDate(nil)

	slog.SetLogLoggerLevel(slog.LevelDebug)

	slog.Info("starting goscores with defaults",
		"gamedate", g,
		"sport", sport,
		"league", league,
	)

	httpClient := http.Client{}
	provider := espn.EspnProvider{Client: &httpClient}

	slog.Debug("fetching scores",
		"provider", provider.GetName(),
	)
	m := provider.GetScores(sport, league, g)
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
