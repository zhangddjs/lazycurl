package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/component"
)

const (
	defaultTime = time.Minute
)

func main() {
	p := tea.NewProgram(component.NewModel(defaultTime))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
