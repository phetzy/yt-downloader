package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// updateURLInput handles updates for the URL input state
func (m *Model) updateURLInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Validate and submit the URL
			url := strings.TrimSpace(m.urlInput.Value())
			if isValidYouTubeURL(url) {
				m.videoURL = url
				m.state = StateLoading
				return m, tea.Batch(
					m.spinner.Tick,
					fetchVideoInfo(url),
				)
			} else if url != "" {
				m.err = ErrInvalidURL
				m.state = StateError
				return m, nil
			}
		case "ctrl+u":
			// Clear input
			m.urlInput.SetValue("")
			return m, nil
		}
	}
	
	// Update the text input
	m.urlInput, cmd = m.urlInput.Update(msg)
	return m, cmd
}

// viewURLInput renders the URL input screen
func (m *Model) viewURLInput() string {
	var b strings.Builder
	
	// Title with ASCII art-style header
	title := `
  ╦ ╦╔═╗╦ ╦╔╦╗╦ ╦╔╗ ╔═╗  ╔╦╗╦ ╦╦  
  ╚╦╝║ ║║ ║ ║ ║ ║╠╩╗║╣    ║ ║ ║║  
   ╩ ╚═╝╚═╝ ╩ ╚═╝╚═╝╚═╝   ╩ ╚═╝╩  
  ╔╦╗╔═╗╦ ╦╔╗╔╦  ╔═╗╔═╗╔╦╗╔═╗╦═╗
   ║║║ ║║║║║║║║  ║ ║╠═╣ ║║║╣ ╠╦╝
  ═╩╝╚═╝╚╩╝╝╚╝╩═╝╚═╝╩ ╩═╩╝╚═╝╩╚═
`
	b.WriteString(RenderTitle(title))
	b.WriteString("\n")
	
	// Subtitle
	b.WriteString(RenderSubtitle("Download YouTube videos with style"))
	b.WriteString("\n\n")
	
	// Instructions
	b.WriteString("Paste your YouTube URL below:\n\n")
	
	// Input box
	b.WriteString(m.urlInput.View())
	b.WriteString("\n\n")
	
	// Help text
	helpText := "Press Enter to continue • Ctrl+U to clear • Ctrl+C to quit"
	b.WriteString(RenderHelp(helpText))
	
	// Center the content
	content := b.String()
	if m.width > 0 {
		content = Center(m.width, content)
	}
	
	return containerStyle.Render(content)
}

// isValidYouTubeURL checks if the URL is a valid YouTube URL
func isValidYouTubeURL(url string) bool {
	url = strings.ToLower(url)
	return strings.Contains(url, "youtube.com/watch") ||
		strings.Contains(url, "youtu.be/") ||
		strings.Contains(url, "youtube.com/v/") ||
		strings.Contains(url, "youtube.com/embed/")
}
