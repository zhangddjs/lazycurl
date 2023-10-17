package request

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/model"
)

type tabState uint

const (
	headerView tabState = iota
	bodyView
	rawView
)

const (
	MODEL_CNT = 3
)

type Model struct {
	state  tabState
	header HeaderModel
	body   BodyModel
	// TODO: auth
	rawCurl string
}

func New(width, height int) Model {
	m := Model{state: headerView}
	m.header = NewHeaderModel(width, height)
	m.body = NewBodyModel(width, height)
	m.rawCurl = ""
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd = m.handleKey(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		cmd = m.handleWindowSize(msg)
		return m, cmd
		// TODO: need update header and body upon analyze success msg coming
	case model.SuccessMsg:
		cmd = m.handleSuccess(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	// TODO: implement the render by active pane
	return ""
}

func (m *Model) handleKey(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg.String() {
	case "ctrl+c", "q":
		return tea.Quit
	case "[":
		m.SwitchToPrevModel()
	case "]":
		m.SwitchToNextModel()
	case "enter":
		// TODO: implement edit popup
	}
	switch m.state {
	// update whichever model is focused
	case headerView:
		m.header, cmd = m.header.Update(msg)
		cmds = append(cmds, cmd)
	case bodyView:
		m.body, cmd = m.body.Update(msg)
		cmds = append(cmds, cmd)
		// TODO: update raw curl view
	}
	return tea.Batch(cmds...)
}

func (m *Model) handleWindowSize(msg tea.WindowSizeMsg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.header, cmd = m.header.Update(msg)
	cmds = append(cmds, cmd)
	m.body, cmd = m.body.Update(msg)
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

func (m *Model) handleSuccess(msg model.SuccessMsg) tea.Cmd {
	var cmd tea.Cmd
	switch msg.Type {
	case model.AnalyzeSuccess:
		data := msg.Data.(model.AnalyzeSuccessData)
		m.header.SetHeader(data.Curl.GetHeader())
		m.body.SetBody(data.Curl.GetBody())
	}
	return cmd
}

func (m Model) isActive(state tabState) bool {
	return m.state == state
}

func (m *Model) SwitchToNextModel() {
	m.state = (m.state + 1) % MODEL_CNT
}

func (m *Model) SwitchToPrevModel() {
	m.state = (m.state - 1) % MODEL_CNT
}
