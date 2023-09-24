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

var strToMethod = map[string]Method{
	"GET":    GET,
	"POST":   POST,
	"PUT":    PUT,
	"DELETE": DELETE,
}

var methodColor = map[Method]string{
	GET:    styles.GREEN,
	POST:   styles.YELLOW,
	PUT:    styles.BLUE,
	DELETE: styles.RED,
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
			return m, SwitchMethod(m.method)
		}
	}
	return m, nil
}

func (m Model) View() string {
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(methodColor[m.method]))
	return helpStyle.Render(m.GetMethod())
}

func (m Model) Render(isActive bool) string {
	if isActive {
		return styles.FocusedMethodStyle.
			SetString(m.View()).
			Render()
	}
	return styles.MethodStyle.SetString(m.View()).Render()
}

func (m *Model) GetMethodColor() string {
	if m == nil {
		return ""
	}
	return methodColor[m.method]
}

func (m *Model) GetMethod() string {
	if m == nil {
		return ""
	}
	return m.method.String()
}

func (m *Model) SetMethod(s string) {
	if method, ok := strToMethod[s]; ok {
		m.method = method
	} else {
		m.method = GET
	}
}

func (m *Model) SwitchToNextMethod() {
	m.method = (m.method + 1) % Method(len(methodStr))
}
