package request

import tea "github.com/charmbracelet/bubbletea"

type BodyModel struct {
	Body string
}

func (m BodyModel) Init() tea.Cmd {
	return nil
}

func (m BodyModel) Update(msg tea.Msg) (BodyModel, tea.Cmd) {
	return m, nil
}

func (m BodyModel) View() string {
	return ""
}
