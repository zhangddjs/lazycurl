package main

import (
	_ "github.com/zhangddjs/lazycurl/conf"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/component"
)

const (
	defaultTime = time.Minute
)

func main() {
	p := tea.NewProgram(
		component.NewModel(defaultTime),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
