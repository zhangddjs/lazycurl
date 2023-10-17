package filemanager

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/zhangddjs/lazycurl/model"
	"github.com/zhangddjs/lazycurl/styles"
)

type BufModel struct {
	Items  []*model.FileNode
	Cursor int
}

func NewBM() BufModel {
	return BufModel{}
}

func (m BufModel) Init() tea.Cmd {
	return nil
}

func (m BufModel) Update(msg tea.Msg) (BufModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd := m.handleKey(msg)
		return m, cmd
	case model.SuccessMsg:
		cmd := m.handleFmSuccess(msg)
		return m, cmd
	}
	return m, nil
}

func (m BufModel) View() string {
	var view strings.Builder
	updateStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(styles.ORANGE))

	for i, item := range m.Items {
		if i == m.Cursor {
			view.WriteString("> ")
		} else {
			view.WriteString("  ")
		}
		// TODO: for imported buffer need a 'N'
		if item.Buffer != item.OriginContent {
			view.WriteString(updateStyle.Render("U "))
			view.WriteString(updateStyle.Render(item.GetName()))
		} else {
			view.WriteString("  ")
			view.WriteString(item.GetName())
		}
		view.WriteString("\n")
	}

	return view.String()
}

func (m BufModel) Render(isActive bool) string {
	if isActive {
		return styles.FocusedFileManagerStyle.Render(m.View())
	}
	return styles.FileManagerStyle.Render(m.View())
}

// GetCurItem get current file item
func (m BufModel) GetCurItem() *model.FileNode {
	if m.Cursor < 0 || m.Cursor >= len(m.Items) {
		return nil
	}
	return m.Items[m.Cursor]
}

// TODO:
// 1.delete Buffer
// 2. import buffer
// 3. save buffer
func (m *BufModel) handleKey(msg tea.KeyMsg) tea.Cmd {
	switch msg.String() {
	case "ctrl+c", "q":
		return tea.Quit
	case "up", "k":
		if m.Cursor > 0 {
			m.Cursor--
		}
	case "down", "j":
		if m.Cursor < len(m.Items)-1 {
			m.Cursor++
		}
	case "enter":
		item := m.GetCurItem()
		if item == nil {
			return nil
		}
		return model.Success(model.OpenBufferSuccess, model.OpenBufferSuccessData{item})
	}
	// Handle keyboard input for navigation and interaction
	// Implement file movement and editing logic here
	return nil
}

// handleFmSuccess
//  1. refresh cursor after user read a file
//  2. append buffer if not exist after user read a file
func (m *BufModel) handleFmSuccess(msg model.SuccessMsg) tea.Cmd {
	switch msg.Type {
	case model.ReadFileSuccess:
		data := msg.Data.(model.ReadFileSuccessData)
		exist := false
		for i, item := range m.Items {
			if data.Item == item {
				m.Cursor = i
				exist = true
			}
		}
		if !exist {
			m.Items = append(m.Items, data.Item)
			m.Cursor = len(m.Items) - 1
		}
	}
	return nil
}
