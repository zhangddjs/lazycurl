package component

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	fm "github.com/zhangddjs/lazycurl/component/filemanager"
	vp "github.com/zhangddjs/lazycurl/component/viewport"
	"github.com/zhangddjs/lazycurl/styles"
)

type sessionState uint

const (
	fmView sessionState = iota
	urlView
	spinnerView
)

const (
	MODEL_CNT   = 3
	defaultTime = time.Minute
)

var (
	// Available spinners
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Points,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
	}
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type mainModel struct {
	state       sessionState
	url         vp.Model
	spinner     spinner.Model
	filemanager fm.Model
	index       int
}

func NewModel(timeout time.Duration) mainModel {
	m := mainModel{state: fmView}
	m.spinner = spinner.New()
	m.filemanager = fm.New()
	m.url = vp.New(44, 1)
	return m
}

func (m mainModel) Init() tea.Cmd {
	// start the timer and spinner on program start
	return tea.Batch(m.spinner.Tick)
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
		case "n":
			if m.state == spinnerView {
				m.Next()
				m.resetSpinner()
				cmds = append(cmds, m.spinner.Tick)
			}
		}
	case tea.WindowSizeMsg:
		m.url, cmd = m.url.Update(msg)
		cmds = append(cmds, cmd)
	}

	switch m.state {
	// update whichever model is focused
	case spinnerView:
		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)
	case fmView:
		m.filemanager, cmd = m.filemanager.Update(msg)
		cmds = append(cmds, cmd)
	case urlView:
		m.url, cmd = m.url.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m mainModel) View() string {
	var s = strings.Builder{}
	model := m.currentFocusedModel()
	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, m.renderLogo(), m.renderMethod(), m.renderUrl()))
	s.WriteString("\n")
	str := lipgloss.JoinVertical(lipgloss.Left, m.renderSpinner(), m.renderSpinner())
	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, m.renderFileManager(), str))
	s.WriteString(helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model)))
	return s.String()
}

func (m mainModel) renderLogo() string {
	return styles.LogoStyle.Render()
}

func (m mainModel) renderMethod() string {
	// TODO: implement me
	return styles.MethodStyle.Render()
}

func (m mainModel) renderUrl() string {
	// TODO: implement me
	if m.state == urlView {
		return styles.FocusedUrlStyle.Render(m.url.View())
	}
	return styles.UrlStyle.Render(m.url.View())
}

func (m mainModel) renderSpinner() string {
	if m.state == spinnerView {
		return styles.FocusedModelStyle.Render(m.spinner.View())
	}
	return styles.ModelStyle.Render(m.spinner.View())
}

func (m mainModel) renderFileManager() string {
	if m.state == fmView {
		return styles.FocusedFileManagerStyle.Render(m.filemanager.View())
	}
	return styles.FileManagerStyle.Render(m.filemanager.View())
}

func (m mainModel) currentFocusedModel() string {
	return "spinner"
}

func (m *mainModel) SwitchToNextModel() {
	m.state = (m.state + 1) % MODEL_CNT
}

func (m *mainModel) Next() {
	if m.index == len(spinners)-1 {
		m.index = 0
	} else {
		m.index++
	}
}

func (m *mainModel) resetSpinner() {
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinners[m.index]
}
