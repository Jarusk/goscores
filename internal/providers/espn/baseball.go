package espn

type baseballScoreboard struct {
	Leagues []struct {
		ID           string `json:"id,omitempty"`
		UID          string `json:"uid,omitempty"`
		Name         string `json:"name,omitempty"`
		Abbreviation string `json:"abbreviation,omitempty"`
		MidsizeName  string `json:"midsizeName,omitempty"`
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
			TimeValid             bool `json:"timeValid,omitempty"`
			NeutralSite           bool `json:"neutralSite,omitempty"`
			ConferenceCompetition bool `json:"conferenceCompetition,omitempty"`
			PlayByPlayAvailable   bool `json:"playByPlayAvailable,omitempty"`
			Recent                bool `json:"recent,omitempty"`
			WasSuspended          bool `json:"wasSuspended,omitempty"`
			Venue                 struct {
				ID       string `json:"id,omitempty"`
				FullName string `json:"fullName,omitempty"`
				Address  struct {
					City  string `json:"city,omitempty"`
					State string `json:"state,omitempty"`
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
					Links            []struct {
						Rel        []string `json:"rel,omitempty"`
						Href       string   `json:"href,omitempty"`
						Text       string   `json:"text,omitempty"`
						IsExternal bool     `json:"isExternal,omitempty"`
						IsPremium  bool     `json:"isPremium,omitempty"`
					} `json:"links,omitempty"`
					Logo string `json:"logo,omitempty"`
				} `json:"team,omitempty"`
				Score      string `json:"score,omitempty"`
				Linescores []struct {
					Value        float64 `json:"value,omitempty"`
					DisplayValue string  `json:"displayValue,omitempty"`
					Period       int     `json:"period,omitempty"`
				} `json:"linescores,omitempty"`
				Statistics []struct {
					Name         string `json:"name,omitempty"`
					Abbreviation string `json:"abbreviation,omitempty"`
					DisplayValue string `json:"displayValue,omitempty"`
				} `json:"statistics,omitempty"`
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
					Statistics []struct {
						Name             string `json:"name,omitempty"`
						Abbreviation     string `json:"abbreviation,omitempty"`
						DisplayValue     string `json:"displayValue,omitempty"`
						RankDisplayValue string `json:"rankDisplayValue,omitempty"`
					} `json:"statistics,omitempty"`
					Record string `json:"record,omitempty"`
				} `json:"probables,omitempty"`
				Hits    int    `json:"hits,omitempty"`
				Errors  int    `json:"errors,omitempty"`
				Record  string `json:"record,omitempty"`
				Records []struct {
					Name         string `json:"name,omitempty"`
					Abbreviation string `json:"abbreviation,omitempty"`
					Type         string `json:"type,omitempty"`
					Summary      string `json:"summary,omitempty"`
				} `json:"records,omitempty"`
			} `json:"competitors,omitempty"`
			Notes []struct {
				Type     string `json:"type,omitempty"`
				Headline string `json:"headline,omitempty"`
			} `json:"notes,omitempty"`
			Situation struct {
				LastPlay struct {
					ID   string `json:"id,omitempty"`
					Type struct {
						ID              string `json:"id,omitempty"`
						Text            string `json:"text,omitempty"`
						Abbreviation    string `json:"abbreviation,omitempty"`
						AlternativeText string `json:"alternativeText,omitempty"`
						Type            string `json:"type,omitempty"`
					} `json:"type,omitempty"`
					Text       string `json:"text,omitempty"`
					ScoreValue int    `json:"scoreValue,omitempty"`
					Team       struct {
						ID string `json:"id,omitempty"`
					} `json:"team,omitempty"`
					AtBatID          string `json:"atBatId,omitempty"`
					SummaryType      string `json:"summaryType,omitempty"`
					AthletesInvolved []struct {
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
					} `json:"athletesInvolved,omitempty"`
				} `json:"lastPlay,omitempty"`
				Balls   int `json:"balls,omitempty"`
				Strikes int `json:"strikes,omitempty"`
				Outs    int `json:"outs,omitempty"`
				Pitcher struct {
					PlayerID int `json:"playerId,omitempty"`
					Period   int `json:"period,omitempty"`
					Athlete  struct {
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
					Summary string `json:"summary,omitempty"`
				} `json:"pitcher,omitempty"`
				Batter struct {
					PlayerID int `json:"playerId,omitempty"`
					Period   int `json:"period,omitempty"`
					Athlete  struct {
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
					Summary string `json:"summary,omitempty"`
				} `json:"batter,omitempty"`
				OnFirst  bool `json:"onFirst,omitempty"`
				OnSecond bool `json:"onSecond,omitempty"`
				OnThird  bool `json:"onThird,omitempty"`
			} `json:"situation,omitempty"`
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
			Format struct {
				Regulation struct {
					Periods int `json:"periods,omitempty"`
				} `json:"regulation,omitempty"`
			} `json:"format,omitempty"`
			StartDate string `json:"startDate,omitempty"`
			Series    struct {
				Type              string `json:"type,omitempty"`
				Title             string `json:"title,omitempty"`
				Summary           string `json:"summary,omitempty"`
				Completed         bool   `json:"completed,omitempty"`
				TotalCompetitions int    `json:"totalCompetitions,omitempty"`
				Competitors       []struct {
					ID   string `json:"id,omitempty"`
					UID  string `json:"uid,omitempty"`
					Wins int    `json:"wins,omitempty"`
					Ties int    `json:"ties,omitempty"`
					Href string `json:"href,omitempty"`
				} `json:"competitors,omitempty"`
			} `json:"series,omitempty"`
			OutsText      string `json:"outsText,omitempty"`
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
		Weather struct {
			DisplayValue    string `json:"displayValue,omitempty"`
			Temperature     int    `json:"temperature,omitempty"`
			HighTemperature int    `json:"highTemperature,omitempty"`
			ConditionID     string `json:"conditionId,omitempty"`
			Link            struct {
				Language   string   `json:"language,omitempty"`
				Rel        []string `json:"rel,omitempty"`
				Href       string   `json:"href,omitempty"`
				Text       string   `json:"text,omitempty"`
				ShortText  string   `json:"shortText,omitempty"`
				IsExternal bool     `json:"isExternal,omitempty"`
				IsPremium  bool     `json:"isPremium,omitempty"`
			} `json:"link,omitempty"`
		} `json:"weather,omitempty"`
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
