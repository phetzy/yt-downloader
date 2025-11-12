package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// updateQualitySelect handles updates for the quality selection state
func (m *Model) updateQualitySelect(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Format selected, move to directory picker
			m.state = StateDirectoryPicker
			return m, nil
		case "up", "k":
			// Move selection up (to be implemented with list component)
			return m, nil
		case "down", "j":
			// Move selection down (to be implemented with list component)
			return m, nil
		}
	}
	
	return m, nil
}

// viewQualitySelect renders the quality selection screen
func (m *Model) viewQualitySelect() string {
	var b strings.Builder
	
	b.WriteString("\n")
	b.WriteString(RenderTitle("🎬 Select Quality"))
	b.WriteString("\n\n")
	
	// Display video information if available
	if info, ok := m.videoInfo.(videoInfoMsg); ok {
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
		b.WriteString("\n\n")
		
		// Display available formats
		b.WriteString("Available Formats:\n\n")
		
		// Video + Audio formats
		hasVideoFormats := false
		for _, format := range info.Formats {
			if format.HasVideo && format.HasAudio {
				if !hasVideoFormats {
					b.WriteString(RenderSubtitle("📹 Video + Audio"))
					b.WriteString("\n")
					hasVideoFormats = true
				}
				size := formatBytes(format.FileSize)
				b.WriteString(fmt.Sprintf("  • %s (%s) - %s - %s\n", 
					format.Quality, 
					format.Resolution,
					format.Format,
					size,
				))
			}
		}
		
		// Audio-only formats
		b.WriteString("\n")
		hasAudioFormats := false
		for _, format := range info.Formats {
			if format.IsAudioOnly {
				if !hasAudioFormats {
					b.WriteString(RenderSubtitle("🎵 Audio Only"))
					b.WriteString("\n")
					hasAudioFormats = true
				}
				size := formatBytes(format.FileSize)
				b.WriteString(fmt.Sprintf("  • %s - %s - %s\n", 
					format.Quality,
					format.Format,
					size,
				))
			}
		}
	}
	
	b.WriteString("\n")
	helpText := "↑/↓ or j/k to navigate • Enter to select • Esc to go back"
	b.WriteString(RenderHelp(helpText))
	
	content := b.String()
	if m.width > 0 {
		content = Center(m.width, content)
	}
	
	return containerStyle.Render(content)
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
