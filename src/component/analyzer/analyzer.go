package analyzer

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/component/filemanager/model"
)

type Curl struct {
	Header  []string
	Body    string
	Rawcurl string
}

type Model struct {
	Cache map[*model.FileNode]Curl
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return ""
}
