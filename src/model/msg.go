package model

import (
	tea "github.com/charmbracelet/bubbletea"
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
	OpenBufferSuccess
	AnalyzeSuccess
)

const (
	UndefinedError errorType = iota
	SaveFileError
	ReadFileError
	CreateFileError
	AnalyzeError
)

type ReadFileSuccessData struct {
	Item *FileNode
}

type OpenBufferSuccessData struct {
	Item *FileNode
}

type AnalyzeSuccessData struct {
	Curl *Curl
}

type SuccessMsg struct {
	Type successType
	Data interface{}
}

type ErrorMsg struct {
	Type errorType
	Msg  string
}

type AnalyzeMsg struct {
	Item *FileNode
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

func Analyze(item *FileNode) tea.Cmd {
	return func() tea.Msg {
		return AnalyzeMsg{item}
	}
}
