package today

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/abdulmuminakinde/futbol/internal/token"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	statusCode   int
	err          error
	response     ApiResponse
	waitingInput bool
}

func InitialModel() model {
	return model{
		statusCode:   0,
		err:          nil,
		response:     ApiResponse{},
		waitingInput: false,
	}
}

type statusMsg int

type errMsg struct {
	err error
}

func (e errMsg) Error() string {
	return e.err.Error()
}

func getData(token string) tea.Cmd {
	return func() tea.Msg {
		var result ApiResponse

		today := time.Now().Format("2006-01-02")
		tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
		url := fmt.Sprintf(
			"https://api.football-data.org/v4/matches?dateFrom=%s&dateTo=%s",
			today,
			tomorrow,
		)
		c := http.Client{Timeout: 10 * time.Second}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return errMsg{err}
		}
		req.Header.Set("X-Auth-Token", token)

		resp, err := c.Do(req)
		if err != nil {
			return errMsg{err}
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return statusMsg(resp.StatusCode)
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return errMsg{err}
		}

		if len(result.Matches) == 0 {
			return tea.Msg(responseMsg(result))
		}

		return tea.Msg(responseMsg(result))
	}
}

type responseMsg ApiResponse

func (m model) Init() tea.Cmd {
	token := token.GetToken()
	return getData(token)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case errMsg:
		m.err = msg.err
		m.waitingInput = true
		return m, nil
	case responseMsg:
		m.response = ApiResponse(msg)
		m.waitingInput = true

		return m, nil
	case tea.KeyMsg:
		if m.waitingInput {
			switch msg.Type {
			case tea.KeyEsc, tea.KeyCtrlC:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	if !m.waitingInput {
		return "Loading..."
	}
	return fmt.Sprint(m.response)
}

func (r ApiResponse) String() string {
	today := time.Now().Format("2006-01-02")
	var result strings.Builder

	result.WriteString(fmt.Sprintf("Matches for %s\n\n", today))
	for _, m := range r.Matches {
		var homeScore, awayScore string
		if m.Score.FullTime.Home != nil {
			score, ok := m.Score.FullTime.Home.(float64)
			if ok {
				homeScore = fmt.Sprintf("%d", int(score))
			} else {
				homeScore = " "
			}
		}
		if m.Score.FullTime.Away != nil {
			score, ok := m.Score.FullTime.Away.(float64)
			if ok {
				awayScore = fmt.Sprintf("%d", int(score))
			} else {
				awayScore = " "
			}
		}
		result.WriteString(fmt.Sprintf(
			"%-25s [%s - %s] %-25s (%s, %s)\n\n",
			m.HomeTeam.Name,
			homeScore, awayScore,
			m.AwayTeam.Name,
			m.Competition.Name,
			m.Status, // "finished", "scheduled", "in_play"
		))
	}
	return result.String()
}
