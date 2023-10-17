package filemanager

import (
	tea "github.com/charmbracelet/bubbletea"
	flags "github.com/jessevdk/go-flags"
	sw "github.com/mattn/go-shellwords"
	"github.com/zhangddjs/lazycurl/model"
	"strings"
)

type AnalyzerModel struct {
}

func NewAnalyzer() AnalyzerModel {
	return AnalyzerModel{}
}

func (m AnalyzerModel) Init() tea.Cmd {
	return nil
}

func (m AnalyzerModel) Update(msg tea.Msg) (AnalyzerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case model.AnalyzeMsg:
		cmd := m.handleAnalyze(msg)
		return m, cmd
	}
	return m, nil
}

func (m AnalyzerModel) View() string {
	return ""
}

func (m *AnalyzerModel) handleAnalyze(msg model.AnalyzeMsg) tea.Cmd {
	data := msg.Item.GetBuffer()
	curl, err := m.analyze(data)
	if err != nil {
		return model.Error(model.AnalyzeError, err.Error())
	}
	return model.Success(model.AnalyzeSuccess, model.AnalyzeSuccessData{&curl})
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
	if len(args) > 1 && args[1] != "" && res.Url == "" {
		res.Url = args[1]
	}
	return res, nil
}
