package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/phetzy/yt-downloader/internal/utils"
)

// updateDirectoryPicker handles updates for the directory picker state
func (m *Model) updateDirectoryPicker(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Initialize directory picker on first entry
	if m.currentDir == "" {
		defaultDir, err := utils.GetDefaultDownloadDir()
		if err != nil {
			m.err = err
			m.state = StateError
			return m, nil
		}
		m.currentDir = defaultDir
		m.loadDirectories()
	}
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// If a directory is selected, enter it
			if m.selectedDirIdx >= 0 && m.selectedDirIdx < len(m.directories) {
				selectedDir := m.directories[m.selectedDirIdx]
				if selectedDir == ".." {
					// Go to parent directory
					m.currentDir = utils.GetParentDir(m.currentDir)
					m.selectedDirIdx = 0
					m.loadDirectories()
				} else if selectedDir == "[SELECT THIS DIRECTORY]" {
					// User selected current directory, proceed to download
					m.downloadPath = m.currentDir
					m.state = StateDownloading
					return m, startDownload()
				} else {
					// Enter the selected subdirectory
					m.currentDir = utils.JoinPath(m.currentDir, selectedDir)
					m.selectedDirIdx = 0
					m.loadDirectories()
				}
			}
			return m, nil
			
		case " ":
			// Space bar selects current directory
			m.downloadPath = m.currentDir
			m.state = StateDownloading
			return m, startDownload()
			
		case "up", "k":
			// Move selection up
			if m.selectedDirIdx > 0 {
				m.selectedDirIdx--
			}
			return m, nil
			
		case "down", "j":
			// Move selection down
			if m.selectedDirIdx < len(m.directories)-1 {
				m.selectedDirIdx++
			}
			return m, nil
			
		case "left", "backspace", "h":
			// Go to parent directory
			m.currentDir = utils.GetParentDir(m.currentDir)
			m.selectedDirIdx = 0
			m.loadDirectories()
			return m, nil
			
		case "right", "l":
			// Enter selected directory
			if m.selectedDirIdx >= 0 && m.selectedDirIdx < len(m.directories) {
				selectedDir := m.directories[m.selectedDirIdx]
				if selectedDir != ".." && selectedDir != "[SELECT THIS DIRECTORY]" {
					m.currentDir = utils.JoinPath(m.currentDir, selectedDir)
					m.selectedDirIdx = 0
					m.loadDirectories()
				}
			}
			return m, nil
		}
	}
	
	return m, nil
}

// loadDirectories loads the list of directories in the current directory
func (m *Model) loadDirectories() {
	dirs, err := utils.ListDirectories(m.currentDir)
	if err != nil {
		// If we can't read the directory, go to parent
		m.currentDir = utils.GetParentDir(m.currentDir)
		dirs, _ = utils.ListDirectories(m.currentDir)
	}
	
	// Build directory list
	m.directories = []string{"[SELECT THIS DIRECTORY]"}
	
	// Add parent directory option if not at root
	if m.currentDir != "/" && m.currentDir != "" {
		m.directories = append(m.directories, "..")
	}
	
	// Add subdirectories
	for _, dir := range dirs {
		if !utils.IsHidden(dir) {
			m.directories = append(m.directories, dir)
		}
	}
}

// viewDirectoryPicker renders the directory picker screen
func (m *Model) viewDirectoryPicker() string {
	var b strings.Builder
	
	b.WriteString("\n")
	b.WriteString(RenderTitle("📁 Choose Download Location"))
	b.WriteString("\n\n")
	
	// Current path
	currentPath := m.currentDir
	if currentPath == "" {
		defaultDir, _ := utils.GetDefaultDownloadDir()
		currentPath = defaultDir
	}
	
	b.WriteString("Current Directory:\n")
	b.WriteString(RenderBox(currentPath, true))
	b.WriteString("\n\n")
	
	// Directory listing
	b.WriteString("Directories:\n\n")
	
	// Show directories with selection indicator
	for i, dir := range m.directories {
		if i == m.selectedDirIdx {
			// Highlight selected directory
			if dir == "[SELECT THIS DIRECTORY]" {
				b.WriteString(selectedItemStyle.Render(fmt.Sprintf("✅ %s", dir)))
			} else if dir == ".." {
				b.WriteString(selectedItemStyle.Render("⬆️  .."))
			} else {
				b.WriteString(selectedItemStyle.Render(fmt.Sprintf("📂 %s", dir)))
			}
		} else {
			// Normal directory
			if dir == "[SELECT THIS DIRECTORY]" {
				b.WriteString(normalItemStyle.Render(fmt.Sprintf("✅ %s", dir)))
			} else if dir == ".." {
				b.WriteString(normalItemStyle.Render("⬆️  .."))
			} else {
				b.WriteString(normalItemStyle.Render(fmt.Sprintf("📂 %s", dir)))
			}
		}
		b.WriteString("\n")
	}
	
	b.WriteString("\n")
	helpText := "↑/↓ or j/k to navigate • Enter to select/enter • Space to choose current • ←/h for parent • Esc to go back"
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
