package espn

import (
	"fmt"

	"github.com/jarusk/goscore/internal/providers"
)

func getScoreboardURL(sport providers.Sport, league providers.League) string {
	return fmt.Sprintf("https://site.api.espn.com/apis/site/v2/sports/%s/%s/scoreboard", sport, league)
}

// Auto-generated with json-to-go
type scoreboard struct {
	Leagues []struct {
		ID           string `json:"id,omitempty"`
		UID          string `json:"uid,omitempty"`
		Name         string `json:"name,omitempty"`
		Abbreviation string `json:"abbreviation,omitempty"`
		Slug         string `json:"slug,omitempty"`
		Season       struct {
			Year        int    `json:"year,omitempty"`
			StartDate   string `json:"startDate,omitempty"`
			EndDate     string `json:"endDate,omitempty"`
			DisplayName string `json:"displayName,omitempty"`
			Type        struct {
				ID           string `json:"id,omitempty"`
				Type         int    `json:"type,omitempty"`
				Name         string `json:"name,omitempty"`
				Abbreviation string `json:"abbreviation,omitempty"`
			} `json:"type,omitempty"`
		} `json:"season,omitempty"`
		Logos []struct {
			Href        string   `json:"href,omitempty"`
			Width       int      `json:"width,omitempty"`
			Height      int      `json:"height,omitempty"`
			Alt         string   `json:"alt,omitempty"`
			Rel         []string `json:"rel,omitempty"`
			LastUpdated string   `json:"lastUpdated,omitempty"`
		} `json:"logos,omitempty"`
		CalendarType        string   `json:"calendarType,omitempty"`
		CalendarIsWhitelist bool     `json:"calendarIsWhitelist,omitempty"`
		CalendarStartDate   string   `json:"calendarStartDate,omitempty"`
		CalendarEndDate     string   `json:"calendarEndDate,omitempty"`
		Calendar            []string `json:"calendar,omitempty"`
	} `json:"leagues,omitempty"`
	Season struct {
		Type int `json:"type,omitempty"`
		Year int `json:"year,omitempty"`
	} `json:"season,omitempty"`
	Day struct {
		Date string `json:"date,omitempty"`
	} `json:"day,omitempty"`
	Events []struct {
		ID        string `json:"id,omitempty"`
		UID       string `json:"uid,omitempty"`
		Date      string `json:"date,omitempty"`
		Name      string `json:"name,omitempty"`
		ShortName string `json:"shortName,omitempty"`
		Season    struct {
			Year int    `json:"year,omitempty"`
			Type int    `json:"type,omitempty"`
			Slug string `json:"slug,omitempty"`
		} `json:"season,omitempty"`
		Competitions []struct {
			ID         string `json:"id,omitempty"`
			UID        string `json:"uid,omitempty"`
			Date       string `json:"date,omitempty"`
			Attendance int    `json:"attendance,omitempty"`
			Type       struct {
				ID           string `json:"id,omitempty"`
				Abbreviation string `json:"abbreviation,omitempty"`
			} `json:"type,omitempty"`
			TimeValid           bool `json:"timeValid,omitempty"`
			NeutralSite         bool `json:"neutralSite,omitempty"`
			PlayByPlayAvailable bool `json:"playByPlayAvailable,omitempty"`
			Recent              bool `json:"recent,omitempty"`
			Venue               struct {
				ID       string `json:"id,omitempty"`
				FullName string `json:"fullName,omitempty"`
				Address  struct {
					City    string `json:"city,omitempty"`
					State   string `json:"state,omitempty"`
					Country string `json:"country,omitempty"`
				} `json:"address,omitempty"`
				Indoor bool `json:"indoor,omitempty"`
			} `json:"venue,omitempty"`
			Competitors []struct {
				ID       string `json:"id,omitempty"`
				UID      string `json:"uid,omitempty"`
				Type     string `json:"type,omitempty"`
				Order    int    `json:"order,omitempty"`
				HomeAway string `json:"homeAway,omitempty"`
				Team     struct {
					ID               string `json:"id,omitempty"`
					UID              string `json:"uid,omitempty"`
					Location         string `json:"location,omitempty"`
					Name             string `json:"name,omitempty"`
					Abbreviation     string `json:"abbreviation,omitempty"`
					DisplayName      string `json:"displayName,omitempty"`
					ShortDisplayName string `json:"shortDisplayName,omitempty"`
					Color            string `json:"color,omitempty"`
					AlternateColor   string `json:"alternateColor,omitempty"`
					IsActive         bool   `json:"isActive,omitempty"`
					Venue            struct {
						ID string `json:"id,omitempty"`
					} `json:"venue,omitempty"`
					Links []struct {
						Rel        []string `json:"rel,omitempty"`
						Href       string   `json:"href,omitempty"`
						Text       string   `json:"text,omitempty"`
						IsExternal bool     `json:"isExternal,omitempty"`
						IsPremium  bool     `json:"isPremium,omitempty"`
					} `json:"links,omitempty"`
					Logo string `json:"logo,omitempty"`
				} `json:"team,omitempty"`
				Score     string `json:"score,omitempty"`
				Probables []struct {
					Name             string `json:"name,omitempty"`
					DisplayName      string `json:"displayName,omitempty"`
					ShortDisplayName string `json:"shortDisplayName,omitempty"`
					Abbreviation     string `json:"abbreviation,omitempty"`
					PlayerID         int    `json:"playerId,omitempty"`
					Athlete          struct {
						ID          string `json:"id,omitempty"`
						FullName    string `json:"fullName,omitempty"`
						DisplayName string `json:"displayName,omitempty"`
						ShortName   string `json:"shortName,omitempty"`
						Links       []struct {
							Rel  []string `json:"rel,omitempty"`
							Href string   `json:"href,omitempty"`
						} `json:"links,omitempty"`
						Headshot string `json:"headshot,omitempty"`
						Jersey   string `json:"jersey,omitempty"`
						Position string `json:"position,omitempty"`
						Team     struct {
							ID string `json:"id,omitempty"`
						} `json:"team,omitempty"`
					} `json:"athlete,omitempty"`
					Status struct {
						ID           string `json:"id,omitempty"`
						Name         string `json:"name,omitempty"`
						Type         string `json:"type,omitempty"`
						Abbreviation string `json:"abbreviation,omitempty"`
					} `json:"status,omitempty"`
					Statistics []any  `json:"statistics,omitempty"`
					Record     string `json:"record,omitempty"`
				} `json:"probables,omitempty"`
				CuratedRank struct {
					Current int `json:"current,omitempty"`
				} `json:"curatedRank,omitempty"`
				Statistics []struct {
					Name             string `json:"name,omitempty"`
					Abbreviation     string `json:"abbreviation,omitempty"`
					DisplayValue     string `json:"displayValue,omitempty"`
					RankDisplayValue string `json:"rankDisplayValue,omitempty"`
				} `json:"statistics,omitempty"`
				Records []struct {
					Name         string `json:"name,omitempty"`
					Abbreviation string `json:"abbreviation,omitempty"`
					Type         string `json:"type,omitempty"`
					Summary      string `json:"summary,omitempty"`
				} `json:"records,omitempty"`
				Leaders []struct {
					Name             string `json:"name,omitempty"`
					DisplayName      string `json:"displayName,omitempty"`
					ShortDisplayName string `json:"shortDisplayName,omitempty"`
					Abbreviation     string `json:"abbreviation,omitempty"`
					Leaders          []struct {
						DisplayValue string  `json:"displayValue,omitempty"`
						Value        float64 `json:"value,omitempty"`
						Athlete      struct {
							ID          string `json:"id,omitempty"`
							FullName    string `json:"fullName,omitempty"`
							DisplayName string `json:"displayName,omitempty"`
							ShortName   string `json:"shortName,omitempty"`
							Links       []struct {
								Rel  []string `json:"rel,omitempty"`
								Href string   `json:"href,omitempty"`
							} `json:"links,omitempty"`
							Headshot string `json:"headshot,omitempty"`
							Jersey   string `json:"jersey,omitempty"`
							Position struct {
								Abbreviation string `json:"abbreviation,omitempty"`
							} `json:"position,omitempty"`
							Team struct {
								ID string `json:"id,omitempty"`
							} `json:"team,omitempty"`
							Active bool `json:"active,omitempty"`
						} `json:"athlete,omitempty"`
						Team struct {
							ID string `json:"id,omitempty"`
						} `json:"team,omitempty"`
					} `json:"leaders,omitempty"`
				} `json:"leaders,omitempty"`
			} `json:"competitors,omitempty"`
			Notes  []any `json:"notes,omitempty"`
			Status struct {
				Clock        float64 `json:"clock,omitempty"`
				DisplayClock string  `json:"displayClock,omitempty"`
				Period       int     `json:"period,omitempty"`
				Type         struct {
					ID          string `json:"id,omitempty"`
					Name        string `json:"name,omitempty"`
					State       string `json:"state,omitempty"`
					Completed   bool   `json:"completed,omitempty"`
					Description string `json:"description,omitempty"`
					Detail      string `json:"detail,omitempty"`
					ShortDetail string `json:"shortDetail,omitempty"`
				} `json:"type,omitempty"`
			} `json:"status,omitempty"`
			Broadcasts []struct {
				Market string   `json:"market,omitempty"`
				Names  []string `json:"names,omitempty"`
			} `json:"broadcasts,omitempty"`
			Format struct {
				Regulation struct {
					Periods int `json:"periods,omitempty"`
				} `json:"regulation,omitempty"`
			} `json:"format,omitempty"`
			Tickets []struct {
				Summary         string `json:"summary,omitempty"`
				NumberAvailable int    `json:"numberAvailable,omitempty"`
				Links           []struct {
					Href string `json:"href,omitempty"`
				} `json:"links,omitempty"`
			} `json:"tickets,omitempty"`
			StartDate     string `json:"startDate,omitempty"`
			Broadcast     string `json:"broadcast,omitempty"`
			GeoBroadcasts []struct {
				Type struct {
					ID        string `json:"id,omitempty"`
					ShortName string `json:"shortName,omitempty"`
				} `json:"type,omitempty"`
				Market struct {
					ID   string `json:"id,omitempty"`
					Type string `json:"type,omitempty"`
				} `json:"market,omitempty"`
				Media struct {
					ShortName string `json:"shortName,omitempty"`
				} `json:"media,omitempty"`
				Lang   string `json:"lang,omitempty"`
				Region string `json:"region,omitempty"`
			} `json:"geoBroadcasts,omitempty"`
			Odds []struct {
				Provider struct {
					ID       string `json:"id,omitempty"`
					Name     string `json:"name,omitempty"`
					Priority int    `json:"priority,omitempty"`
				} `json:"provider,omitempty"`
				Details      string  `json:"details,omitempty"`
				OverUnder    float64 `json:"overUnder,omitempty"`
				Spread       float64 `json:"spread,omitempty"`
				AwayTeamOdds struct {
					Favorite bool `json:"favorite,omitempty"`
					Underdog bool `json:"underdog,omitempty"`
					Team     struct {
						ID           string `json:"id,omitempty"`
						UID          string `json:"uid,omitempty"`
						Abbreviation string `json:"abbreviation,omitempty"`
						Name         string `json:"name,omitempty"`
						DisplayName  string `json:"displayName,omitempty"`
						Logo         string `json:"logo,omitempty"`
					} `json:"team,omitempty"`
				} `json:"awayTeamOdds,omitempty"`
				HomeTeamOdds struct {
					Favorite bool `json:"favorite,omitempty"`
					Underdog bool `json:"underdog,omitempty"`
					Team     struct {
						ID           string `json:"id,omitempty"`
						UID          string `json:"uid,omitempty"`
						Abbreviation string `json:"abbreviation,omitempty"`
						Name         string `json:"name,omitempty"`
						DisplayName  string `json:"displayName,omitempty"`
						Logo         string `json:"logo,omitempty"`
					} `json:"team,omitempty"`
				} `json:"homeTeamOdds,omitempty"`
				Open struct {
					Over struct {
						Value                 float64 `json:"value,omitempty"`
						DisplayValue          string  `json:"displayValue,omitempty"`
						AlternateDisplayValue string  `json:"alternateDisplayValue,omitempty"`
						Decimal               float64 `json:"decimal,omitempty"`
						Fraction              string  `json:"fraction,omitempty"`
						American              string  `json:"american,omitempty"`
					} `json:"over,omitempty"`
					Under struct {
						Value                 float64 `json:"value,omitempty"`
						DisplayValue          string  `json:"displayValue,omitempty"`
						AlternateDisplayValue string  `json:"alternateDisplayValue,omitempty"`
						Decimal               float64 `json:"decimal,omitempty"`
						Fraction              string  `json:"fraction,omitempty"`
						American              string  `json:"american,omitempty"`
					} `json:"under,omitempty"`
					Total struct {
						Value                 float64 `json:"value,omitempty"`
						DisplayValue          string  `json:"displayValue,omitempty"`
						AlternateDisplayValue string  `json:"alternateDisplayValue,omitempty"`
						Decimal               float64 `json:"decimal,omitempty"`
						Fraction              string  `json:"fraction,omitempty"`
						American              string  `json:"american,omitempty"`
					} `json:"total,omitempty"`
				} `json:"open,omitempty"`
				Current struct {
					Over struct {
						Value                 float64 `json:"value,omitempty"`
						DisplayValue          string  `json:"displayValue,omitempty"`
						AlternateDisplayValue string  `json:"alternateDisplayValue,omitempty"`
						Decimal               float64 `json:"decimal,omitempty"`
						Fraction              string  `json:"fraction,omitempty"`
						American              string  `json:"american,omitempty"`
						Outcome               struct {
							Type string `json:"type,omitempty"`
						} `json:"outcome,omitempty"`
					} `json:"over,omitempty"`
					Under struct {
						Value                 float64 `json:"value,omitempty"`
						DisplayValue          string  `json:"displayValue,omitempty"`
						AlternateDisplayValue string  `json:"alternateDisplayValue,omitempty"`
						Decimal               float64 `json:"decimal,omitempty"`
						Fraction              string  `json:"fraction,omitempty"`
						American              string  `json:"american,omitempty"`
						Outcome               struct {
							Type string `json:"type,omitempty"`
						} `json:"outcome,omitempty"`
					} `json:"under,omitempty"`
					Total struct {
						AlternateDisplayValue string `json:"alternateDisplayValue,omitempty"`
						American              string `json:"american,omitempty"`
					} `json:"total,omitempty"`
				} `json:"current,omitempty"`
			} `json:"odds,omitempty"`
			Highlights []any `json:"highlights,omitempty"`
		} `json:"competitions,omitempty"`
		Links []struct {
			Language   string   `json:"language,omitempty"`
			Rel        []string `json:"rel,omitempty"`
			Href       string   `json:"href,omitempty"`
			Text       string   `json:"text,omitempty"`
			ShortText  string   `json:"shortText,omitempty"`
			IsExternal bool     `json:"isExternal,omitempty"`
			IsPremium  bool     `json:"isPremium,omitempty"`
		} `json:"links,omitempty"`
		Status struct {
			Clock        float64 `json:"clock,omitempty"`
			DisplayClock string  `json:"displayClock,omitempty"`
			Period       int     `json:"period,omitempty"`
			Type         struct {
				ID          string `json:"id,omitempty"`
				Name        string `json:"name,omitempty"`
				State       string `json:"state,omitempty"`
				Completed   bool   `json:"completed,omitempty"`
				Description string `json:"description,omitempty"`
				Detail      string `json:"detail,omitempty"`
				ShortDetail string `json:"shortDetail,omitempty"`
			} `json:"type,omitempty"`
		} `json:"status,omitempty"`
	} `json:"events,omitempty"`
}
