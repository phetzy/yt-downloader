package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// updateComplete handles updates for the completion state
func (m *Model) updateComplete(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "n":
			// Download another video
			m.state = StateURLInput
			m.err = nil
			m.videoURL = ""
			m.videoInfo = nil
			m.selectedFormat = nil
			m.downloadPath = ""
			m.urlInput.SetValue("")
			m.urlInput.Focus()
			return m, nil
		case "o":
			// Open folder (to be implemented)
			return m, nil
		}
	}
	
	return m, nil
}

// viewComplete renders the completion screen
func (m *Model) viewComplete() string {
	var b strings.Builder
	
	b.WriteString("\n\n")
	b.WriteString(RenderSuccess("✅ Download Complete!"))
	b.WriteString("\n\n")
	
	// Display download location
	if m.downloadPath != "" {
		b.WriteString("File saved to:\n")
		b.WriteString(RenderBox(m.downloadPath, false))
	} else {
		b.WriteString("File saved successfully!\n")
	}
	
	b.WriteString("\n\n")
	
	// Options
	b.WriteString("What would you like to do?\n\n")
	b.WriteString("  • Press Enter or N to download another video\n")
	b.WriteString("  • Press O to open download folder\n")
	b.WriteString("  • Press Q or Ctrl+C to quit\n")
	
	b.WriteString("\n")
	helpText := "Thank you for using YouTube TUI Downloader!"
	b.WriteString(RenderHelp(helpText))
	
	content := b.String()
	if m.width > 0 {
		content = Center(m.width, content)
	}
	
	return containerStyle.Render(content)
}
