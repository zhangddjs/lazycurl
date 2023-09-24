package httpmethod

import tea "github.com/charmbracelet/bubbletea"

type SwitchMethodMsg struct {
	Method Method
}

func SwitchMethod(m Method) tea.Cmd {
	return func() tea.Msg {
		return SwitchMethodMsg{m}
	}
}
