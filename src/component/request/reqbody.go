package request

import (
	tea "github.com/charmbracelet/bubbletea"
	vp "github.com/zhangddjs/lazycurl/component/viewport"
	"github.com/zhangddjs/lazycurl/model"
	"github.com/zhangddjs/lazycurl/styles"
)

type BodyModel struct {
	body     string
	viewport vp.Model
}

func NewBodyModel(width, height int) BodyModel {
	m := BodyModel{}
	m.viewport = vp.New(width, height, "")
	return m
}

func (m BodyModel) Init() tea.Cmd {
	return nil
}

func (m BodyModel) Update(msg tea.Msg) (BodyModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd = m.handleKey(msg)
		return m, cmd
	case tea.WindowSizeMsg:
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	case model.SuccessMsg:
		cmd = m.handleSuccess(msg)
		return m, cmd
	}
	return m, nil
}

func (m BodyModel) View() string {
	return m.viewport.View()
}

func (m BodyModel) Render(isActive bool) string {
	if isActive {
		return styles.FocusedRespBodyStyle.Render(m.View())
	}
	return styles.RespBodyStyle.Render(m.View())
}

func (m BodyModel) GetBody() string {
	return m.body
}

func (m *BodyModel) SetBody(body string) {
	if m == nil {
		return
	}
	m.body = body
	m.viewport.SetContent(body)
}

func (m *BodyModel) handleKey(msg tea.KeyMsg) tea.Cmd {
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

func (m *BodyModel) handleSuccess(msg model.SuccessMsg) tea.Cmd {
	var cmd tea.Cmd
	switch msg.Type {
	case model.AnalyzeSuccess:
		data := msg.Data.(model.AnalyzeSuccessData)
		m.SetBody(data.Curl.GetBody())
	}
	return cmd
}
