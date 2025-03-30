package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/jarusk/goscore/internal/providers/espn"
)

func main() {
	defdate := time.Now()
	gamedate := fmt.Sprintf("%d%d%d", defdate.Year(), defdate.Month(), defdate.Day())

	slog.Info("starting goscores with defaults",
		"gamedate", gamedate,
		"sport", "hockey",
		"league", "nhl",
	)

	httpClient := http.Client{}
	provider := espn.EspnProvider{Client: &httpClient}

	matches := provider.GetScores("hockey", "nhl", time.Now())
	fmt.Printf("%+v", matches)
}
