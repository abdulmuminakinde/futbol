package today

import "time"

type ApiResponse struct {
	Matches []struct {
		UtcDate     time.Time `json:"utcDate"`
		Status      string    `json:"status"`
		Matchday    int       `json:"matchday"`
		Competition struct {
			Name string `json:"name"`
		} `json:"competition"`
		HomeTeam struct {
			Name string `json:"name"`
		} `json:"homeTeam"`
		AwayTeam struct {
			Name string `json:"name"`
		} `json:"awayTeam"`
		Score struct {
			FullTime struct {
				Home interface{} `json:"home"`
				Away interface{} `json:"away"`
			} `json:"fullTime"`
		} `json:"score"`
	} `json:"matches"`
}
