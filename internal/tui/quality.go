package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// updateQualitySelect handles updates for the quality selection state
func (m *Model) updateQualitySelect(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.qualityList.SetWidth(msg.Width)
		m.qualityList.SetHeight(msg.Height - 10)
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Get selected item
			if i, ok := m.qualityList.SelectedItem().(qualityItem); ok {
				m.selectedFormat = i.format
				m.state = StateDirectoryPicker
				return m, nil
			}
		}
	}
	
	// Update the list component
	var cmd tea.Cmd
	m.qualityList, cmd = m.qualityList.Update(msg)
	return m, cmd
}

// viewQualitySelect renders the quality selection screen
func (m *Model) viewQualitySelect() string {
	var b strings.Builder
	
	b.WriteString("\n")
	b.WriteString(RenderTitle("🎬 Select Quality"))
	b.WriteString("\n")
	
	// Display video information if available
	if info, ok := m.videoInfo.(videoInfoMsg); ok {
		// Update list items if not already set
		if len(m.qualityList.Items()) == 0 && len(info.Formats) > 0 {
			items := convertFormatsToItems(info.Formats)
			m.qualityList.SetItems(items)
		}
		
		b.WriteString(RenderBox(fmt.Sprintf(
			"Title:    %s\n"+
			"Channel:  %s\n"+
			"Duration: %s\n"+
			"Views:    %s\n"+
			"Uploaded: %s",
			info.Title,
			info.Author,
			info.Duration,
			info.Views,
			info.UploadDate,
		), false))
		b.WriteString("\n")
	}
	
	// Render the list component
	b.WriteString(m.qualityList.View())
	b.WriteString("\n")
	
	helpText := "↑/↓ or j/k to navigate • Enter to select • Esc to go back • q to quit"
	b.WriteString(RenderHelp(helpText))
	
	return b.String()
}

// formatBytes converts bytes to human-readable format
func formatBytes(bytes int64) string {
	const (
		KB = 1024
		MB = 1024 * KB
		GB = 1024 * MB
	)
	
	switch {
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/float64(GB))
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/float64(MB))
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/float64(KB))
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}
