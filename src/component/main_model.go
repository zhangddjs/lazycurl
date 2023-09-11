package component

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	fm "github.com/zhangddjs/lazycurl/component/filemanager"
	hm "github.com/zhangddjs/lazycurl/component/httpmethod"
	vp "github.com/zhangddjs/lazycurl/component/viewport"
	"github.com/zhangddjs/lazycurl/styles"
)

type sessionState uint

var (
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

const (
	fmView sessionState = iota
	methodView
	urlView
)

const (
	MODEL_CNT   = 3
	defaultTime = time.Minute
)

type mainModel struct {
	state       sessionState
	method      hm.Model
	url         vp.Model
	filemanager fm.Model
	index       int
}

func NewModel(timeout time.Duration) mainModel {
	m := mainModel{state: fmView}
	m.filemanager = fm.New()
	m.url = vp.New(44, 1)
	return m
}

func (m mainModel) Init() tea.Cmd {
	// start the timer and spinner on program start
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			m.SwitchToNextModel()
			// TODO: case "s" save buffer to file
		}
		switch m.state {
		// update whichever model is focused
		case fmView:
			m.filemanager, cmd = m.filemanager.Update(msg)
			cmds = append(cmds, cmd)
		case urlView:
			m.url, cmd = m.url.Update(msg)
			cmds = append(cmds, cmd)
		case methodView:
			m.method, cmd = m.method.Update(msg)
			cmds = append(cmds, cmd)
		}
	case tea.WindowSizeMsg:
		m.url, cmd = m.url.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s = strings.Builder{}
	model := m.currentFocusedModel()

	fm := m.filemanager.Render(m.isActive(fmView))
	logo := styles.LogoStyle.Render()
	method := m.method.Render(m.isActive(methodView))
	url := m.url.RenderUrl(m.isActive(urlView))

	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, logo, method, url))
	s.WriteString("\n")
	str := lipgloss.JoinVertical(lipgloss.Left, "  ", "  ")
	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, fm, str))
	s.WriteString(helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model)))
	return s.String()
}

func (m mainModel) currentFocusedModel() string {
	return "spinner"
}

func (m mainModel) isActive(state sessionState) bool {
	return m.state == state
}

func (m *mainModel) SwitchToNextModel() {
	m.state = (m.state + 1) % MODEL_CNT
}
