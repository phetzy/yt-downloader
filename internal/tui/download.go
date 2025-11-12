package tui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// updateDownloading handles updates for the downloading state
func (m *Model) updateDownloading(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case downloadProgressMsg:
		// Update progress (to be implemented with actual progress bar)
		// For now, simulate completion after a moment
		if msg.BytesDownloaded >= msg.TotalBytes {
			m.state = StateComplete
			return m, nil
		}
		return m, nil
		
	case downloadCompleteMsg:
		m.downloadPath = msg.FilePath
		m.state = StateComplete
		return m, nil
		
	case errMsg:
		m.err = msg.err
		m.state = StateError
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			// Cancel download
			m.quitting = true
			return m, tea.Quit
		}
	}
	
	return m, nil
}

// viewDownloading renders the download progress screen
func (m *Model) viewDownloading() string {
	var b strings.Builder
	
	b.WriteString("\n")
	b.WriteString(RenderTitle("⬇️  Downloading..."))
	b.WriteString("\n\n")
	
	// Progress information (placeholder)
	b.WriteString("Progress: 45%\n\n")
	b.WriteString("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n")
	
	b.WriteString(fmt.Sprintf("Downloaded: %s / %s\n", "22.5 MB", "50.0 MB"))
	b.WriteString(fmt.Sprintf("Speed:      %s\n", "5.2 MB/s"))
	b.WriteString(fmt.Sprintf("ETA:        %s\n", "5 seconds"))
	
	b.WriteString("\n")
	helpText := "Ctrl+C to cancel download"
	b.WriteString(RenderHelp(helpText))
	
	content := b.String()
	if m.width > 0 {
		content = Center(m.width, content)
	}
	
	return containerStyle.Render(content)
}

// formatDuration formats a duration in seconds to human-readable format
func formatDuration(seconds int) string {
	duration := time.Duration(seconds) * time.Second
	
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60
	
	if hours > 0 {
		return fmt.Sprintf("%dh %dm %ds", hours, minutes, secs)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm %ds", minutes, secs)
	}
	return fmt.Sprintf("%ds", secs)
}

// formatSpeed formats download speed in bytes/sec to human-readable format
func formatSpeed(bytesPerSec float64) string {
	const (
		KB = 1024.0
		MB = 1024.0 * KB
		GB = 1024.0 * MB
	)
	
	switch {
	case bytesPerSec >= GB:
		return fmt.Sprintf("%.2f GB/s", bytesPerSec/GB)
	case bytesPerSec >= MB:
		return fmt.Sprintf("%.2f MB/s", bytesPerSec/MB)
	case bytesPerSec >= KB:
		return fmt.Sprintf("%.2f KB/s", bytesPerSec/KB)
	default:
		return fmt.Sprintf("%.0f B/s", bytesPerSec)
	}
}
