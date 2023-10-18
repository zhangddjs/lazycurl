package request

import (
	tea "github.com/charmbracelet/bubbletea"
	vp "github.com/zhangddjs/lazycurl/component/viewport"
	"github.com/zhangddjs/lazycurl/styles"
)

type RawModel struct {
	raw      string
	viewport vp.Model
}

func NewRawModel(width, height int) RawModel {
	m := RawModel{}
	m.viewport = vp.New(width, height, "")
	return m
}

func (m RawModel) Init() tea.Cmd {
	return nil
}

func (m RawModel) Update(msg tea.Msg) (RawModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd = m.handleKey(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m RawModel) View() string {
	return m.viewport.View()
}

func (m RawModel) Render(isActive bool) string {
	if isActive {
		return styles.FocusedRespBodyStyle.Render(m.View())
	}
	return styles.RespBodyStyle.Render(m.View())
}

func (m RawModel) GetRawcurl() string {
	return m.raw
}

func (m *RawModel) SetRawcurl(raw string) {
	if m == nil {
		return
	}
	m.raw = raw
	m.viewport.SetContent(raw)
}

func (m *RawModel) handleKey(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg.String() {
	case "ctrl+c", "q":
		return tea.Quit
	case "enter":
		// TODO: implement edit popup
	}
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}
