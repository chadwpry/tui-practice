package models

import (
	"log"
	"tuipractice/styles"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type WelcomeKeys struct{}

func (k WelcomeKeys) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k WelcomeKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{},
	}
}

var WelcomeKeyMap = WelcomeKeys{}

type WelcomeModel struct {
	width  int
	height int
}

func NewWelcome() WelcomeModel {
	return WelcomeModel{}
}

func (m WelcomeModel) Init() tea.Cmd {
	return nil
}

func (m WelcomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		default:
			log.Printf("Welcome.Update(%v)::KeyMsg", msg)
		}
	default:
		log.Printf("Welcome.Update(%v)::default", msg)
	}

	return m, nil
}

func (m WelcomeModel) View() string {
	log.Printf("Welcome.View() w: %d h: %d", m.width, m.height)

	return styles.SectionStyle("Welcome")
}
