package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// updateLoading handles updates for the loading state
func (m *Model) updateLoading(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case videoInfoMsg:
		// Video info fetched successfully
		m.videoInfo = msg
		m.state = StateQualitySelect
		return m, nil
		
	case errMsg:
		// Error occurred while fetching video info
		m.err = msg.err
		m.state = StateError
		return m, nil
	}
	
	// Update the spinner
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

// viewLoading renders the loading screen
func (m *Model) viewLoading() string {
	var b strings.Builder
	
	b.WriteString("\n\n")
	b.WriteString(m.spinner.View())
	b.WriteString("  Fetching video information...\n\n")
	b.WriteString(RenderHelp("This may take a moment"))
	
	content := b.String()
	if m.width > 0 {
		content = Center(m.width, content)
	}
	
	return containerStyle.Render(content)
}
