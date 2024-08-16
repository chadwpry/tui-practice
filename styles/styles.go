package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("190")).Render
	SectionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("140")).Render
)
