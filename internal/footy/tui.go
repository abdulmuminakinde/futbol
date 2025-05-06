package footy

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/abdulmuminakinde/futbol/internal/token"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	statusCode   int
	response     Response
	err          error
	waitingInput bool
	done         bool
}

type responseMsg Response

type statusMsg int

type errMsg struct {
	err error
}

func (e errMsg) Error() string {
	return e.err.Error()
}

func getData(token string) tea.Cmd {
	return func() tea.Msg {
		var response Response

		url := "https://api.football-data.org/v4/matches"
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

		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return errMsg{err}
		}

		return responseMsg(response)
	}
}

func InitialModel() model {
	return model{
		statusCode:   0,
		response:     Response{},
		err:          nil,
		waitingInput: false,
		done:         false,
	}
}

func (m model) Init() tea.Cmd {
	token := token.GetToken()
	return getData(token)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case responseMsg:
		m.response = Response(msg)
		return m, nil
	case errMsg:
		m.err = msg.err
		return m, nil
	case tea.KeyMsg:
		if m.waitingInput {
			switch msg.Type {
			case tea.KeyEsc, tea.KeyCtrlC:
				m.waitingInput = false
				m.done = true
				return m, nil
			}
		} else {
			switch msg.Type {
			case tea.KeyEsc, tea.KeyCtrlC:
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.done {
		return "Done!\n\nPress Esc to quit"
	}
	if m.err != nil {
		return m.err.Error()
	}
	return fmt.Sprint(m.response)
}
