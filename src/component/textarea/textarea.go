package textarea

// A simple program demonstrating the textarea component from the Bubbles
// component library.

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/styles"
)

type errMsg error

type Model struct {
	textarea textarea.Model
	err      error
}

func New(w, h int) Model {
	ti := textarea.New()
	ti.Placeholder = "Once upon a time..."
	ti.Focus()
	ti.SetHeight(h)
	ti.SetWidth(w)
	// TODO: max height, max width

	return Model{
		textarea: ti,
		err:      nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.textarea.View()
}

func (m Model) RenderReqBody(isActive bool) string {
	if isActive {
		return styles.FocusedTextAreaStyle.Render(m.View())
	}
	return styles.TextAreaStyle.Render(m.View())
}
