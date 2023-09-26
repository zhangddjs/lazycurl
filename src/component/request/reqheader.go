package request

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zhangddjs/lazycurl/styles"
)

type HeaderModel struct {
	Headers []string
	Cursor  int
}

func (m HeaderModel) Init() tea.Cmd {
	return nil
}

func (m HeaderModel) Update(msg tea.Msg) (HeaderModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd := m.handleKey(msg)
		return m, cmd
	}
	return m, nil
}

func (m HeaderModel) View() string {
	var view strings.Builder
	style := lipgloss.NewStyle().Background(lipgloss.Color(styles.GREEN))
	for i, header := range m.Headers {
		if i == m.Cursor {
			view.WriteString(style.Render(header))
		} else {
			view.WriteString(header)
		}
		if i < len(m.Headers) {
			view.WriteString("\n")
		}
	}
	return view.String()
}

// GetCurItem get current file item
func (m HeaderModel) GetCurHeader() string {
	if m.Cursor < 0 || m.Cursor >= len(m.Headers) {
		return ""
	}
	return m.Headers[m.Cursor]
}

func (m *HeaderModel) handleKey(msg tea.KeyMsg) tea.Cmd {
	switch msg.String() {
	case "ctrl+c", "q":
		return tea.Quit
	case "up", "k":
		if m.Cursor > 0 {
			m.Cursor--
		}
	case "down", "j":
		if m.Cursor < len(m.Headers)-1 {
			m.Cursor++
		}
	case "enter":
		header := m.GetCurHeader()
		_ = header
		// TODO: implement edit header popup
	}
	return nil
}
