package filemanager

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/component/filemanager/model"
)

type (
	successType uint
	errorType   uint
)

const (
	UndefinedSuccess successType = iota
	SaveFileSuccess
	ReadFileSuccess
	CreateFileSuccess
)

const (
	UndefinedError errorType = iota
	SaveFileError
	ReadFileError
	CreateFileError
)

type ReadFileSuccessData struct {
	Item *model.FileNode
}

type SuccessMsg struct {
	Type successType
	data interface{}
}

type ErrorMsg struct {
	Type errorType
	Msg  string
}

type AnalyzeMsg struct {
	Content string
}

func Success(t successType, data interface{}) tea.Cmd {
	return func() tea.Msg {
		return SuccessMsg{t, data}
	}
}

func Error(t errorType, msg string) tea.Cmd {
	return func() tea.Msg {
		return ErrorMsg{t, msg}
	}
}

func Analyze(content string) tea.Cmd {
	return func() tea.Msg {
		return AnalyzeMsg{content}
	}
}
