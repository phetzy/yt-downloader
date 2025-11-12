package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// updateError handles updates for the error state
func (m *Model) updateError(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "r":
			// Retry - go back to URL input
			m.state = StateURLInput
			m.err = nil
			m.urlInput.SetValue("")
			m.urlInput.Focus()
			return m, nil
		}
	}
	
	return m, nil
}

// viewError renders the error screen
func (m *Model) viewError() string {
	var b strings.Builder
	
	b.WriteString("\n\n")
	b.WriteString(RenderError("❌ Error"))
	b.WriteString("\n\n")
	
	// Display the error message
	if m.err != nil {
		b.WriteString(m.err.Error())
	} else {
		b.WriteString("An unknown error occurred")
	}
	
	b.WriteString("\n\n")
	
	// Help text
	helpText := "Press Enter or R to retry • Ctrl+C to quit"
	b.WriteString(RenderHelp(helpText))
	
	content := b.String()
	if m.width > 0 {
		content = Center(m.width, content)
	}
	
	return containerStyle.Render(content)
}
