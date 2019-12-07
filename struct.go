package main

import "time"

type standing struct {
	Filters struct {
	} `json:"filters"`
	Competition struct {
		ID   int `json:"id"`
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"area"`
		Name        string    `json:"name"`
		Code        string    `json:"code"`
		Plan        string    `json:"plan"`
		LastUpdated time.Time `json:"lastUpdated"`
	} `json:"competition"`
	Season struct {
		ID              int         `json:"id"`
		StartDate       string      `json:"startDate"`
		EndDate         string      `json:"endDate"`
		CurrentMatchday int         `json:"currentMatchday"`
		Winner          interface{} `json:"winner"`
	} `json:"season"`
	Standings []struct {
		Stage string      `json:"stage"`
		Type  string      `json:"type"`
		Group interface{} `json:"group"`
		Table []struct {
			Position int `json:"position"`
			Team     struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				CrestURL string `json:"crestUrl"`
			} `json:"team"`
			PlayedGames    int `json:"playedGames"`
			Won            int `json:"won"`
			Draw           int `json:"draw"`
			Lost           int `json:"lost"`
			Points         int `json:"points"`
			GoalsFor       int `json:"goalsFor"`
			GoalsAgainst   int `json:"goalsAgainst"`
			GoalDifference int `json:"goalDifference"`
		} `json:"table"`
	} `json:"standings"`
}

type team struct {
	Name          string
	Points        int
	AveragePoints float32
}
