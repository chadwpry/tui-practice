package models

import (
	"log"
	"tuipractice/styles"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RouterKeys struct {
	Esc      key.Binding
	Quit     key.Binding
	Pagefile key.Binding
	Welcome  key.Binding
}

func (k RouterKeys) ShortHelp() []key.Binding {
	return []key.Binding{k.Esc, k.Pagefile, k.Welcome, k.Quit}
}

func (k RouterKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Esc, k.Pagefile, k.Welcome, k.Quit},
	}
}

var RouterKeyMap = RouterKeys{
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "return"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),
	Pagefile: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "page file"),
	),
	Welcome: key.NewBinding(
		key.WithKeys("w"),
		key.WithHelp("w", "welcome"),
	),
}

type RouterModel struct {
	keys     RouterKeys
	help     help.Model
	modelKey string
	models   map[string]tea.Model
	width    int
	height   int
}

const (
	Pagefile string = "pagefile"
	Welcome  string = "welcome"
)

func NewRouter() RouterModel {
	return RouterModel{
		keys:     RouterKeyMap,
		modelKey: Welcome,
		models: map[string]tea.Model{
			Welcome:  NewWelcome(),
			Pagefile: NewPagefile("docs/artichoke.md"),
		},
		help: help.New(),
	}
}

func (m RouterModel) Init() tea.Cmd {
	return nil
}

func (m RouterModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = msg.Width

		m.models[m.modelKey], cmd = m.models[m.modelKey].Update(msg)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, RouterKeyMap.Quit):
			return m, tea.Quit
		case key.Matches(msg, RouterKeyMap.Pagefile):
			m.modelKey = Pagefile
			m.models[m.modelKey], cmd = m.models[m.modelKey].Update(tea.WindowSizeMsg{
				Width:  m.width,
				Height: m.height,
			})
		case key.Matches(msg, RouterKeyMap.Welcome):
			m.modelKey = Welcome
			m.models[m.modelKey], cmd = m.models[m.modelKey].Update(tea.WindowSizeMsg{
				Width:  m.width,
				Height: m.height,
			})
		default:
			log.Printf("Router.Update(%v)::KeyMsg", msg)

			m.models[m.modelKey], cmd = m.models[m.modelKey].Update(msg)
		}
	default:
		log.Printf("Router.Update(%v)::msg.(type)", msg)
	}

	return m, cmd
}

func (m RouterModel) View() string {
	log.Printf("Router.View() w: %d h: %d", m.width, m.height)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(
			lipgloss.Left,
			styles.TitleStyle("Router"),
			m.models[m.modelKey].View(),
			m.help.View(m.keys),
		),
	)
}
