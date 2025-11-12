package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// updateDirectoryPicker handles updates for the directory picker state
func (m *Model) updateDirectoryPicker(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			// Directory selected, start download
			m.state = StateDownloading
			return m, startDownload()
		case "up", "k":
			// Move selection up (to be implemented)
			return m, nil
		case "down", "j":
			// Move selection down (to be implemented)
			return m, nil
		case "left", "backspace", "h":
			// Go to parent directory (to be implemented)
			return m, nil
		case "right", "l":
			// Enter selected directory (to be implemented)
			return m, nil
		}
	}
	
	return m, nil
}

// viewDirectoryPicker renders the directory picker screen
func (m *Model) viewDirectoryPicker() string {
	var b strings.Builder
	
	b.WriteString("\n")
	b.WriteString(RenderTitle("📁 Choose Download Location"))
	b.WriteString("\n\n")
	
	// Current path
	currentPath := m.downloadPath
	if currentPath == "" {
		currentPath = "~/Downloads" // Default
	}
	
	b.WriteString("Current Directory:\n")
	b.WriteString(RenderBox(currentPath, true))
	b.WriteString("\n\n")
	
	// Directory listing (placeholder)
	b.WriteString("Directories:\n\n")
	b.WriteString("  • ..\n")
	b.WriteString("  • Documents\n")
	b.WriteString("  • Downloads\n")
	b.WriteString("  • Videos\n")
	b.WriteString("  • Music\n")
	
	b.WriteString("\n")
	helpText := "↑/↓ to navigate • Enter/Space to select • ←/→ to change directory • Esc to go back"
	b.WriteString(RenderHelp(helpText))
	
	content := b.String()
	if m.width > 0 {
		content = Center(m.width, content)
	}
	
	return containerStyle.Render(content)
}

// startDownload initiates the download process
func startDownload() tea.Cmd {
	return func() tea.Msg {
		// TODO: Implement actual download
		// For now, simulate progress updates
		return downloadProgressMsg{
			BytesDownloaded: 0,
			TotalBytes:      50000000,
			Speed:           0,
			ETA:             0,
		}
	}
}
