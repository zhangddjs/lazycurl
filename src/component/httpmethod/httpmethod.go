package httpmethod

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zhangddjs/lazycurl/styles"
)

type Method uint

const (
	GET Method = iota
	POST
	PUT
	DELETE
)

var methodStr = [...]string{
	"GET", "POST", "PUT", "DELETE",
}

var methodColor = map[Method]string{
	GET:    "#32CD32", //Green
	POST:   "#FFFF00", //yellow
	PUT:    "#1E90FF", //Blue
	DELETE: "#FF0000", //Red
}

func (m Method) String() string {
	if m < 0 || m >= Method(len(methodStr)) {
		return "Unknown"
	}
	return methodStr[m]
}

type Model struct {
	method Method
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "n":
			m.SwitchToNextMethod()
			// TODO: implement cmd to update raw file
			return m, nil
		}
		// TODO: case open file, read the method
	}
	return m, nil
}

func (m Model) View() string {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(methodColor[m.method]))
	return helpStyle.Render(m.method.String())
}

func (m Model) Render(isActive bool) string {
	if isActive {
		return styles.FocusedMethodStyle.
			SetString(m.View()).
			Render()
	}
	return styles.MethodStyle.SetString(m.View()).Render()
}

func (m *Model) SwitchToNextMethod() {
	m.method = (m.method + 1) % Method(len(methodStr))
}
