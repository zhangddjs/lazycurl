package analyzer

import (
	tea "github.com/charmbracelet/bubbletea"
)

type (
	successType uint
	errorType   uint
)

const (
	UndefinedSuccess successType = iota
	AnalyzeSuccess
)

const (
	UndefinedError errorType = iota
	AnalyzeError
)

type AnalyzeSuccessData struct {
	Curl *Curl
}

type SuccessMsg struct {
	Type successType
	data interface{}
}

type ErrorMsg struct {
	Type errorType
	Msg  string
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
