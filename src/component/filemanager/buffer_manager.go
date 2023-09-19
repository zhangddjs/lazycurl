package filemanager

import (
	"errors"
	"os"
	"path/filepath"
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
	return BufModel{}
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
			data.Item.Name = "testtest.curl"

		}
	}
	return m, nil
}

func (m BufModel) View() string {
	var view strings.Builder

	// view.WriteString("Root Directory: " + m.BasePath + "\n\n")

	for i, item := range m.Items {
		if i == m.Cursor {
			view.WriteString("> ")
		} else {
			view.WriteString("  ")
		}
		if item.Buffer != item.OriginContent {
			view.WriteString("U ")
		}
		view.WriteString(item.GetName())
		if item.IsDir() {
			view.WriteString("/")
		}
		view.WriteString("\n")
	}

	return view.String()
}

func (m BufModel) GetCurItem() *model.FileNode {
	if m.Cursor < 0 || m.Cursor >= len(m.Items) {
		return nil
	}
	return m.Items[m.Cursor]
}
