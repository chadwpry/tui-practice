package main

import (
	"log"
	"os"
	"tuipractice/models"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	f, err := tea.LogToFile("logs/dev.log", "tui-practice")
	if err != nil {
		log.Fatalf("Issue opening log file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(models.NewRouter(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
