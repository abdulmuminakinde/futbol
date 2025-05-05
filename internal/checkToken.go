package footy

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	token        string
	err          error
	tokenInput   textinput.Model
	waitingInput bool
	done         bool
}

func getConfigDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/.config/futbol", home)
}

func getFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/.config/futbol/token.txt", home)
}

type tokenMsg string

type errMsg struct {
	err error
}

func (e errMsg) Error() string { return e.err.Error() }

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Token"
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 20
	return model{
		tokenInput:   ti,
		err:          nil,
		waitingInput: false,
		done:         false,
	}
}

func setToken(token string) tea.Cmd {
	return func() tea.Msg {
		filePath := getFilePath()
		if err := os.MkdirAll(getConfigDir(), 0o755); err != nil {
			return errMsg{err: err}
		}
		f, err := os.Create(filePath)
		if err != nil {
			return errMsg{err: err}
		}
		defer f.Close()
		_, err = f.WriteString(token)
		if err != nil {
			return errMsg{err: err}
		}
		return tokenMsg(token)
	}
}

func checkToken() tea.Cmd {
	return func() tea.Msg {
		filePath := getFilePath()
		_, err := os.Stat(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				return errMsg{err: errors.New("token-required")}
			}
			return errMsg{err: err}
		}

		data, err := os.ReadFile(filePath)
		if err != nil {
			return errMsg{err: err}
		}
		token := strings.TrimSpace(string(data))
		if token == "" {
			return errMsg{err: errors.New("token-required")}
		}

		return tokenMsg(token)
	}
}

func (m model) Init() tea.Cmd {
	return checkToken()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tokenMsg:
		m.token = string(msg)
		return m, nil

	case errMsg:
		if msg.err.Error() == "token-required" {
			m.waitingInput = true
			return m, nil
		}
		m.err = msg.err
		return m, tea.Quit

	case tea.KeyMsg:
		if m.waitingInput {
			switch msg.Type {
			case tea.KeyEnter:
				token := strings.TrimSpace(m.tokenInput.Value())
				if token != "" {
					m.done = true
					return m, setToken(token)
				}
				return m, nil

			case tea.KeyEsc, tea.KeyCtrlC:
				return m, tea.Quit
			default:
				m.tokenInput, cmd = m.tokenInput.Update(msg)
				return m, cmd
			}
		} else {
			switch msg.Type {
			case tea.KeyEnter, tea.KeyEsc, tea.KeyCtrlC:
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
	if m.done {
		return fmt.Sprint("Done!\n\nPress Esc to quit")
	}

	if m.waitingInput {
		return fmt.Sprintf(
			"Please enter your token:\n\n%s\n\nPress Enter to save or Esc to Quit",
			m.tokenInput.View(),
		)
	}

	if m.token != "" {
		return fmt.Sprintf("Token: %s\n\nPress any key to quit", m.token)
	}
	return m.token
}
