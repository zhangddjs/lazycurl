package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/model"
)

const (
	defaultTime = time.Minute
)

func main() {
	p := tea.NewProgram(model.NewModel(defaultTime))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
