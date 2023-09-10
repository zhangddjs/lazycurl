package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zhangddjs/lazycurl/tutorial/filemanager/component"
)

func main() {
	model := component.New()
	p := tea.NewProgram(&model)
	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
