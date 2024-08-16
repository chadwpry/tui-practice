package models

import (
	"fmt"
	"log"
	"os"
	"strings"
	"tuipractice/styles"
	"tuipractice/utils"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()
	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
	useHighPerformanceRenderer = false
)

type PagefileKeys struct{}

func (k PagefileKeys) ShortHelp() []key.Binding {
	return []key.Binding{}
}

func (k PagefileKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{},
	}
}

var PagefileKeyMap = PagefileKeys{}

type PagefileModel struct {
	content  string
	ready    bool
	viewport viewport.Model
	width    int
	height   int
}

func NewPagefile(filepath string) PagefileModel {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("failed to load file %v", err)
		os.Exit(1)
	}

	return PagefileModel{
		content: string(content),
	}
}

func (m PagefileModel) Init() tea.Cmd {
	return nil
}

func (m PagefileModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(m.width, m.height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.content)
			m.ready = true

			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = m.width
			m.viewport.Height = m.height
		}

		if useHighPerformanceRenderer {
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	case tea.KeyMsg:
		switch {
		default:
			log.Printf("Pagefile.Update(%v)::tea.KeyMsg", msg)
		}
	default:
		log.Printf("Pagefile.Update(%v)::msg.(type)", msg)
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m PagefileModel) View() string {
	log.Printf("Pagefile.View() w: %d h: %d", m.width, m.height)

	if !m.ready {
		return styles.SectionStyle("Pagefile Initializing...")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.headerView(),
		m.viewport.View(),
		m.footerView(),
	)
}

func (m PagefileModel) headerView() string {
	title := styles.TitleStyle("Pagefile")
	line := strings.Repeat("-", utils.Max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m PagefileModel) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("-", utils.Max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
