package analyzer

import (
	tea "github.com/charmbracelet/bubbletea"
	flags "github.com/jessevdk/go-flags"
	sw "github.com/mattn/go-shellwords"
	"github.com/zhangddjs/lazycurl/component/filemanager"
	"github.com/zhangddjs/lazycurl/component/filemanager/model"
	"strings"
)

type Model struct {
	Curls map[*model.FileNode]Curl
}

func New() Model {
	return Model{Curls: make(map[*model.FileNode]Curl)}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case filemanager.AnalyzeMsg:
		data := msg.Content
		curl, err := m.analyze(data)
		if err != nil {
			return m, Error(AnalyzeError, err.Error())
		}
		return m, Success(AnalyzeSuccess, AnalyzeSuccessData{&curl})
	}
	return m, nil
}

func (m Model) View() string {
	return ""
}

func (m Model) analyze(data string) (Curl, error) {
	res := Curl{Rawcurl: data}
	cmdParts, err := sw.Parse(data)
	if err != nil {
		return res, err
	}
	for i, str := range cmdParts {
		cmdParts[i] = strings.TrimSpace(str)
	}
	args, _ := flags.ParseArgs(&res, cmdParts)
	// args[0] normally be 'curl'
	if len(args) > 1 && args[1] != "" && res.Url != "" {
		res.Url = args[1]
	}
	return res, nil
}
