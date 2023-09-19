package filemanager

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/component/filemanager/model"
	"github.com/zhangddjs/lazycurl/styles"
)

type BufModel struct {
	Items       []*model.FileNode
	Cursor      int
	ExistingBuf map[*model.FileNode]bool
}

func NewBM() BufModel {
	return BufModel{ExistingBuf: make(map[*model.FileNode]bool)}
}

func (m BufModel) Init() tea.Cmd {
	return nil
}

func (m BufModel) Update(msg tea.Msg) (BufModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Cursor < len(m.Items)-1 {
				m.Cursor++
			}

		case "enter":
			item := m.GetCurItem()
			// TODO: send analyze Cmd
			_ = item

		}
		// Handle keyboard input for navigation and interaction
		// Implement file movement and editing logic here
	case SuccessMsg:
		if msg.Type == ReadFileSuccess {
			data := msg.data.(ReadFileSuccessData)
			if !m.ExistingBuf[data.Item] {
				m.ExistingBuf[data.Item] = true
				m.Items = append(m.Items, data.Item)
			}
		}
	}
	return m, nil
}

func (m BufModel) View() string {
	var view strings.Builder

	for i, item := range m.Items {
		if i == m.Cursor {
			view.WriteString("> ")
		} else {
			view.WriteString("  ")
		}
		if item.Buffer != item.OriginContent {
			view.WriteString("U ")
		} else {
			view.WriteString("  ")
		}
		view.WriteString(item.GetName())
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

func (m BufModel) GetCurItem() *model.FileNode {
	if m.Cursor < 0 || m.Cursor >= len(m.Items) {
		return nil
	}
	return m.Items[m.Cursor]
}
