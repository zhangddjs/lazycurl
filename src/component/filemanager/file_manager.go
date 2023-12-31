package filemanager

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/model"
	"github.com/zhangddjs/lazycurl/styles"
)

var (
	ErrInvalidFileType = errors.New("file type not curl")
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

func New() Model {
	pwd, _ := os.Getwd()
	model := Model{
		ItemsUnderExpandedDir: make(map[*model.FileNode]bool),
		ExpandedDirItems:      make(map[*model.FileNode][]*model.FileNode),
		BasePath:              pwd,
	}
	model.loadRootFiles()
	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd := m.handleKey(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	var view strings.Builder

	// view.WriteString("Root Directory: " + m.BasePath + "\n\n")

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

func (m Model) Render(isActive bool) string {
	if isActive {
		return styles.FocusedFileManagerStyle.Render(m.View())
	}
	return styles.FileManagerStyle.Render(m.View())
}

func (m *Model) handleKey(msg tea.KeyMsg) tea.Cmd {
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
		if item.IsDir() {
			if m.isDirExpanded(item) {
				m.foldSubFiles(item)
			} else {
				m.ExpandedDirItems[item] = m.loadSubFiles(item)
			}
		} else if item.IsCurl() {
			// TODO:implement
			// 1. load file
			// 2. analyze file into request method, params, body, header...
			// 3. if return err then show pop up
			// 4. return cmd to refresh the text area of Request infomation
			if item.GetBuffer() == "" { // TODO: maybe need better judge way
				err := m.readFile()
				if err != nil {
					return model.Error(model.ReadFileError, err.Error())
				}
			}
			return model.Success(model.ReadFileSuccess, model.ReadFileSuccessData{item})
		}
	}
	return nil
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

// readFile read file from disk
func (m *Model) readFile() error {
	// TODO: if already opened then no need to read io again
	item := m.GetCurItem()
	if !item.IsCurl() {
		return ErrInvalidFileType
	}
	path := item.GetFullName()
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	contentStr := string(fileContent)
	item.Buffer = contentStr
	item.OriginContent = contentStr
	// TODO: BufferManager append this item to buffer list

	return nil
}

// TODO:load the file from pwd
// 1. get pwd
// 2. read file into buffer
// 3. send cmd and raw curl to main in purpose to update the other components
// 4. analyze the file(need analyzer to do it), to get：
// 		* request url
// 		* request header
// 		* request body
// 		* request auth
// 		* request method
// 5. send analyzed data to other components
//

// GetCurItem get current file item
func (m Model) GetCurItem() *model.FileNode {
	if m.Cursor < 0 || m.Cursor >= len(m.Items) {
		return nil
	}
	return m.Items[m.Cursor]
}

func (m Model) isDirExpanded(dir *model.FileNode) bool {
	if _, ok := m.ExpandedDirItems[dir]; !ok {
		return false
	}
	return true
}

func (m Model) getFilesInDir(dir *model.FileNode) []*model.FileNode {
	return m.ExpandedDirItems[dir] // Change this according to your logic
}
