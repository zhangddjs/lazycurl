package request

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	vp "github.com/zhangddjs/lazycurl/component/viewport"
	"github.com/zhangddjs/lazycurl/styles"
)

type HeaderModel struct {
	headers  []string
	cursor   int
	viewport vp.Model
}

func NewHeaderModel(width, height int) HeaderModel {
	m := HeaderModel{}
	m.viewport = vp.New(width, height, "")
	return m
}

func (m HeaderModel) Init() tea.Cmd {
	return nil
}

func (m HeaderModel) Update(msg tea.Msg) (HeaderModel, tea.Cmd) {
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

func (m HeaderModel) View() string {
	var view strings.Builder
	style := lipgloss.NewStyle().Background(lipgloss.Color(styles.CHARM))
	for i, header := range m.headers {
		if i == m.cursor {
			view.WriteString(style.Render(header))
		} else {
			view.WriteString(header)
		}
		if i < len(m.headers) {
			view.WriteString("\n")
		}
	}
	m.viewport.SetContent(view.String())
	return m.viewport.View()
}

func (m HeaderModel) Render(isActive bool) string {
	if isActive {
		return styles.FocusedRespBodyStyle.Render(m.View())
	}
	return styles.RespBodyStyle.Render(m.View())
}

// GetCurItem get current file item
func (m HeaderModel) GetCurHeader() string {
	if m.cursor < 0 || m.cursor >= len(m.headers) {
		return ""
	}
	return m.headers[m.cursor]
}

func (m *HeaderModel) SetHeader(headers []string) {
	if m == nil {
		return
	}
	m.headers = headers
	var view strings.Builder
	style := lipgloss.NewStyle().Background(lipgloss.Color(styles.CHARM))
	for i, header := range m.headers {
		if i == m.cursor {
			view.WriteString(style.Render(header))
		} else {
			view.WriteString(header)
		}
		if i < len(m.headers) {
			view.WriteString("\n")
		}
	}
	m.viewport.SetContent(view.String())
}

func (m *HeaderModel) handleKey(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg.String() {
	case "ctrl+c", "q":
		return tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.headers)-1 {
			m.cursor++
		}
		// TODO: add page up page down support
	case "enter":
		header := m.GetCurHeader()
		_ = header
		// TODO: implement edit header popup
	}
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}
