package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/tutorial/test/model"
)

type Model struct {
	Items    []*model.FileNode
	Cursor   int
	BasePath string
	// 目录展开状态的映射，用于跟踪每个目录是否展开
	ItemsUnderExpandedDir map[*model.FileNode]bool
	// 存储每个目录下的文件列表
	ExpandedDirItems map[*model.FileNode][]*model.FileNode
}

func main() {
	pwd, _ := os.Getwd()
	model := Model{
		ItemsUnderExpandedDir: make(map[*model.FileNode]bool),
		ExpandedDirItems:      make(map[*model.FileNode][]*model.FileNode),
		BasePath:              pwd,
	}
	model.loadRootFiles()

	p := tea.NewProgram(&model)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func (m *Model) loadRootFiles() []*model.FileNode {
	m.Items = make([]*model.FileNode, 0)
	path := m.BasePath

	files, _ := filepath.Glob(filepath.Join(path, "*"))

	for _, f := range files {
		base := filepath.Base(f)
		if strings.HasSuffix(base, ".curl") {
			item := &model.FileNode{
				Name:       base,
				Type:       model.FileType_Curl,
				Path:       path,
				ParentNode: nil,
			}
			m.Items = append(m.Items, item)
		} else if fi, err := os.Stat(f); err == nil && fi.IsDir() {
			item := &model.FileNode{
				Name:       base,
				Type:       model.FileType_Dir,
				Path:       path,
				ParentNode: nil,
			}
			m.Items = append(m.Items, item)
		}
	}
	return m.Items
}

func (m *Model) loadSubFiles(dir *model.FileNode) []*model.FileNode {
	if dir == nil {
		return make([]*model.FileNode, 0)
	}

	path := dir.GetFullName()

	files, _ := filepath.Glob(filepath.Join(path, "*"))
	items := make([]*model.FileNode, 0)

	for _, f := range files {
		base := filepath.Base(f)
		if strings.HasSuffix(base, ".curl") {
			item := &model.FileNode{
				Name:       base,
				Type:       model.FileType_Curl,
				Path:       path,
				ParentNode: dir,
				Level:      dir.Level + 1,
			}
			items = append(items, item)
			m.ItemsUnderExpandedDir[item] = true
		} else if fi, err := os.Stat(f); err == nil && fi.IsDir() {
			item := &model.FileNode{
				Name:       base,
				Type:       model.FileType_Dir,
				Path:       path,
				ParentNode: dir,
				Level:      dir.Level + 1,
			}
			items = append(items, item)
			m.ItemsUnderExpandedDir[item] = true
		}
	}
	m.Items = append(m.Items[:m.Cursor+1], append(items, m.Items[m.Cursor+1:]...)...)
	return items
}

func (m *Model) foldSubFiles(dir *model.FileNode) {
	m.markNeedFoldFiles(dir)
	m.doFold(dir)
}

func (m *Model) markNeedFoldFiles(dir *model.FileNode) {
	if dir == nil {
		return
	}
	for _, item := range m.ExpandedDirItems[dir] {
		if m.isDirExpanded(item) {
			m.markNeedFoldFiles(item)
		}
		delete(m.ItemsUnderExpandedDir, item)
	}
	delete(m.ExpandedDirItems, dir)
}

func (m *Model) doFold(dir *model.FileNode) {
	if dir == nil {
		return
	}
	items := make([]*model.FileNode, 0)
	for _, item := range m.Items {
		if !m.ItemsUnderExpandedDir[item] && item.GetLevel() > 0 {
			continue
		}
		items = append(items, item)
	}
	m.Items = items
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.Cursor < len(m.Items) {
				m.Cursor++
			}

		case "enter":
			// 当用户按下 Enter 键时，切换目录的展开状态
			item := m.Items[m.Cursor]
			if item.IsDir() {
				if m.isDirExpanded(item) {
					m.foldSubFiles(item)
				} else {
					m.ExpandedDirItems[item] = m.loadSubFiles(item)
				}
			}

			// 如果目录是展开的，可以在这里获取目录下的文件列表并更新到 m.DirFiles 中
			return m, nil
		}
		// Handle keyboard input for navigation and interaction
		// Implement file movement and editing logic here
	}
	return m, nil
}

func (m *Model) View() string {
	var view strings.Builder

	view.WriteString("Root Directory: " + m.BasePath + "\n\n")

	for i, item := range m.Items {
		if i == m.Cursor {
			view.WriteString("> ")
		} else {
			view.WriteString("  ")
		}
		indent := strings.Repeat("  ", item.GetLevel())
		view.WriteString(indent)
		if item.IsDir() {
			if m.isDirExpanded(item) {
				view.WriteString("- ")
			} else {
				view.WriteString("+ ")
			}
		} else {
			view.WriteString("* ")
		}
		view.WriteString(item.GetName())
		if item.IsDir() {
			view.WriteString("/")
		}
		view.WriteString("\n")
	}

	return view.String()
}

func (m *Model) isDirExpanded(dir *model.FileNode) bool {
	if _, ok := m.ExpandedDirItems[dir]; !ok {
		return false
	}
	return true
}

func (m *Model) getFilesInDir(dir *model.FileNode) []*model.FileNode {
	return m.ExpandedDirItems[dir] // Change this according to your logic
}
