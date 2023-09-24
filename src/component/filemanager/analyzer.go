package filemanager

import (
	tea "github.com/charmbracelet/bubbletea"
	flags "github.com/jessevdk/go-flags"
	sw "github.com/mattn/go-shellwords"
	"github.com/zhangddjs/lazycurl/component/filemanager/model"
	"strings"
)

type AnalyzerModel struct {
	Curls map[*model.FileNode]model.Curl
}

func NewAnalyzer() AnalyzerModel {
	return AnalyzerModel{Curls: make(map[*model.FileNode]model.Curl)}
}

func (m AnalyzerModel) Init() tea.Cmd {
	return nil
}

func (m AnalyzerModel) Update(msg tea.Msg) (AnalyzerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case AnalyzeMsg:
		data := msg.Item.Buffer
		curl, err := m.analyze(data)
		if err != nil {
			return m, Error(AnalyzeError, err.Error())
		}
		return m, Success(AnalyzeSuccess, AnalyzeSuccessData{&curl})
	}
	return m, nil
}

func (m AnalyzerModel) View() string {
	return ""
}

func (m AnalyzerModel) analyze(data string) (model.Curl, error) {
	res := model.Curl{Rawcurl: data}
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
