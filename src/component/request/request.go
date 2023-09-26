package request

import tea "github.com/charmbracelet/bubbletea"

type TabState uint

const (
	HeaderView TabState = iota
	BodyView
	RawView
)

type Model struct {
	State  TabState
	Header HeaderModel
	Body   BodyModel
	// TODO: auth
	RawCurl string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return ""
}
