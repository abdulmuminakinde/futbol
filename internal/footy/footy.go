package footy

import (
	"fmt"
	"time"
)

func (r Response) String() string {
	return fmt.Sprint(r.Matches)
}

type Response struct {
	Filters struct {
		DateFrom   string `json:"dateFrom"`
		DateTo     string `json:"dateTo"`
		Permission string `json:"permission"`
	} `json:"filters"`
	ResultSet struct {
		Count        int    `json:"count"`
		Competitions string `json:"competitions"`
		First        string `json:"first"`
		Last         string `json:"last"`
		Played       int    `json:"played"`
	} `json:"resultSet"`
	Matches []struct {
		Area struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
			Flag string `json:"flag"`
		} `json:"area"`
		Competition struct {
			ID     int    `json:"id"`
			Name   string `json:"name"`
			Code   string `json:"code"`
			Type   string `json:"type"`
			Emblem string `json:"emblem"`
		} `json:"competition"`
		Season struct {
			ID              int         `json:"id"`
			StartDate       string      `json:"startDate"`
			EndDate         string      `json:"endDate"`
			CurrentMatchday int         `json:"currentMatchday"`
			Winner          interface{} `json:"winner"`
		} `json:"season"`
		ID          int         `json:"id"`
		UtcDate     time.Time   `json:"utcDate"`
		Status      string      `json:"status"`
		Matchday    int         `json:"matchday"`
		Stage       string      `json:"stage"`
		Group       interface{} `json:"group"`
		LastUpdated time.Time   `json:"lastUpdated"`
		HomeTeam    struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			ShortName string `json:"shortName"`
			Tla       string `json:"tla"`
			Crest     string `json:"crest"`
		} `json:"homeTeam"`
		AwayTeam struct {
			ID        int    `json:"id"`
			Name      string `json:"name"`
			ShortName string `json:"shortName"`
			Tla       string `json:"tla"`
			Crest     string `json:"crest"`
		} `json:"awayTeam"`
		Score struct {
			Winner   interface{} `json:"winner"`
			Duration string      `json:"duration"`
			FullTime struct {
				Home interface{} `json:"home"`
				Away interface{} `json:"away"`
			} `json:"fullTime"`
			HalfTime struct {
				Home interface{} `json:"home"`
				Away interface{} `json:"away"`
			} `json:"halfTime"`
		} `json:"score"`
		Odds struct {
			Msg string `json:"msg"`
		} `json:"odds"`
		Referees []interface{} `json:"referees"`
	} `json:"matches"`
}
