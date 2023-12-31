package viewport

// An example program demonstrating the pager component from the Bubbles
// component library.

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/styles"
)

// You generally won't need this unless you're processing stuff with
// complicated ANSI escape sequences. Turn it on if you notice flickering.
//
// Also keep in mind that high performance rendering only works for programs
// that use the full size of the terminal. We're enabling that below with
// tea.EnterAltScreen().
const useHighPerformanceRenderer = false

type Model struct {
	content  string
	ready    bool
	viewport viewport.Model
}

func New(width, height int, content string) Model {
	m := Model{content: content}
	m.viewport = viewport.New(width, height)
	m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
	return m
}

func (m *Model) SetContent(content string) {
	m.content = content
	m.viewport.SetContent(content)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	// TODO: need support focused size change
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(m.viewport.Width, m.viewport.Height)
			m.viewport.YPosition = 0
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.content)
			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = 1
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			//
			// This is needed for high-performance rendering only.
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s", m.viewport.View())
}

func (m Model) RenderUrl(isActive bool) string {
	if isActive {
		return styles.FocusedUrlStyle.Render(m.View())
	}
	return styles.UrlStyle.Render(m.View())
}

func (m Model) RenderRespBody(isActive bool) string {
	if isActive {
		return styles.FocusedRespBodyStyle.Render(m.View())
	}
	return styles.RespBodyStyle.Render(m.View())
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
