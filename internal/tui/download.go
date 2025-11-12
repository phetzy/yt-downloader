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
		// Update progress
		if msg.TotalBytes > 0 {
			m.downloadProgress = float64(msg.BytesDownloaded) / float64(msg.TotalBytes)
		}
		
		// Check if download is complete
		if msg.BytesDownloaded >= msg.TotalBytes {
			m.state = StateComplete
			return m, nil
		}
		
		// Continue receiving progress updates
		return m, nil
		
	case downloadCompleteMsg:
		m.downloadPath = msg.FilePath
		m.downloadProgress = 1.0
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
		
	case tea.WindowSizeMsg:
		m.progressBar.Width = msg.Width - 4
		return m, nil
	}
	
	return m, nil
}

// viewDownloading renders the download progress screen
func (m *Model) viewDownloading() string {
	var b strings.Builder
	
	b.WriteString("\n")
	b.WriteString(RenderTitle("⬇️  Downloading..."))
	b.WriteString("\n\n")
	
	// Display video title if available
	if info, ok := m.videoInfo.(videoInfoMsg); ok {
		b.WriteString(fmt.Sprintf("File: %s\n\n", info.Title))
	}
	
	// Render the progress bar
	percentage := m.downloadProgress * 100
	b.WriteString(fmt.Sprintf("Progress: %.1f%%\n\n", percentage))
	b.WriteString(m.progressBar.ViewAs(m.downloadProgress))
	b.WriteString("\n\n")
	
	// Download statistics (placeholder values for now)
	// In production, these would come from the actual download progress message
	b.WriteString(fmt.Sprintf("Downloaded: %.1f MB / %.1f MB\n", 
		m.downloadProgress*50.0, 50.0))
	b.WriteString(fmt.Sprintf("Speed:      %.2f MB/s\n", 5.2))
	eta := int((1.0 - m.downloadProgress) * 10)
	if eta > 0 {
		b.WriteString(fmt.Sprintf("ETA:        %s\n", formatDuration(eta)))
	} else {
		b.WriteString("ETA:        calculating...\n")
	}
	
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
