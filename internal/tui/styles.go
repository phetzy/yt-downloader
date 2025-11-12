package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	// Color palette
	primaryColor = lipgloss.Color("#FF0000") // YouTube red
	accentColor  = lipgloss.Color("#00ADD8") // Go blue
	successColor = lipgloss.Color("#00FF00")
	errorColor   = lipgloss.Color("#FF0000")
	subtleColor  = lipgloss.Color("#888888")
	
	// Text styles
	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Padding(1, 0).
			Align(lipgloss.Center)
	
	subtitleStyle = lipgloss.NewStyle().
			Foreground(accentColor).
			Italic(true).
			Padding(0, 0, 1, 0)
	
	helpStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			Italic(true).
			Padding(1, 0)
	
	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true).
			Padding(1, 2)
	
	successStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)
	
	// Border styles
	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(accentColor).
			Padding(1, 2).
			Margin(1, 0)
	
	activeBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2).
			Margin(1, 0)
	
	// List styles
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(primaryColor).
				Bold(true).
				PaddingLeft(2)
	
	normalItemStyle = lipgloss.NewStyle().
			PaddingLeft(4)
	
	// Spinner style
	spinnerStyle = lipgloss.NewStyle().
			Foreground(accentColor)
	
	// Layout styles
	containerStyle = lipgloss.NewStyle().
			Padding(2, 4)
)

// Helper functions for styling

// RenderTitle renders a styled title
func RenderTitle(text string) string {
	return titleStyle.Render(text)
}

// RenderSubtitle renders a styled subtitle
func RenderSubtitle(text string) string {
	return subtitleStyle.Render(text)
}

// RenderHelp renders styled help text
func RenderHelp(text string) string {
	return helpStyle.Render(text)
}

// RenderError renders styled error text
func RenderError(text string) string {
	return errorStyle.Render(text)
}

// RenderSuccess renders styled success text
func RenderSuccess(text string) string {
	return successStyle.Render(text)
}

// RenderBox renders content in a bordered box
func RenderBox(content string, active bool) string {
	if active {
		return activeBoxStyle.Render(content)
	}
	return boxStyle.Render(content)
}

// Center centers content horizontally
func Center(width int, content string) string {
	return lipgloss.Place(
		width,
		lipgloss.Height(content),
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}
