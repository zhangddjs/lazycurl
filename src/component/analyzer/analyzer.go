package analyzer

import (
	tea "github.com/charmbracelet/bubbletea"
	sw "github.com/mattn/go-shellwords"
	"github.com/zhangddjs/lazycurl/component/filemanager"
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
	switch msg := msg.(type) {
	case filemanager.AnalyzeMsg:
		data := msg.Content
		curl := m.analyze(data)
		// TODO: send analyzed cmd with curl obj
		_ = curl
	}
	return m, nil
}

func (m Model) View() string {
	return ""
}

func (m Model) analyze(data string) Curl {
	s := "curl -X POST -H \"Content-Type: application/json\" -d '{\"key\":\"value\"}' https://example.com/api"
	cmdParts, err := sw.Parse(data)
	if err != nil {
		// TODO: send error msg
		return Curl{}
	}
	_ = cmdParts
	return Curl{}
}
