package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// AppState represents the current state of the application
type AppState int

const (
	StateURLInput AppState = iota
	StateLoading
	StateQualitySelect
	StateDirectoryPicker
	StateDownloading
	StateComplete
	StateError
)

// Model is the main application model for Bubble Tea
type Model struct {
	// Application state
	state   AppState
	err     error
	width   int
	height  int
	
	// Component states
	urlInput    textinput.Model
	spinner     spinner.Model
	
	// Data
	videoURL    string
	videoInfo   interface{} // Will be replaced with actual video info struct
	selectedFormat interface{}
	downloadPath string
	
	// Flags
	quitting    bool
}

// NewApp creates and initializes a new application model
func NewApp() *Model {
	// Initialize URL input
	ti := textinput.New()
	ti.Placeholder = "https://www.youtube.com/watch?v=..."
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 60
	
	// Initialize spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle
	
	return &Model{
		state:    StateURLInput,
		urlInput: ti,
		spinner:  s,
	}
}

// Init initializes the application
func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages and updates the model
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
		
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.state == StateURLInput || m.state == StateComplete || m.state == StateError {
				m.quitting = true
				return m, tea.Quit
			}
		case "esc":
			// Allow going back from certain states
			if m.state == StateQualitySelect {
				m.state = StateURLInput
				m.urlInput.SetValue("")
				m.urlInput.Focus()
				return m, nil
			}
		}
	}
	
	// Delegate to state-specific update handlers
	switch m.state {
	case StateURLInput:
		return m.updateURLInput(msg)
	case StateLoading:
		return m.updateLoading(msg)
	case StateQualitySelect:
		return m.updateQualitySelect(msg)
	case StateDirectoryPicker:
		return m.updateDirectoryPicker(msg)
	case StateDownloading:
		return m.updateDownloading(msg)
	case StateComplete:
		return m.updateComplete(msg)
	case StateError:
		return m.updateError(msg)
	}
	
	return m, nil
}

// View renders the application
func (m *Model) View() string {
	if m.quitting {
		return ""
	}
	
	// Delegate to state-specific view renderers
	switch m.state {
	case StateURLInput:
		return m.viewURLInput()
	case StateLoading:
		return m.viewLoading()
	case StateQualitySelect:
		return m.viewQualitySelect()
	case StateDirectoryPicker:
		return m.viewDirectoryPicker()
	case StateDownloading:
		return m.viewDownloading()
	case StateComplete:
		return m.viewComplete()
	case StateError:
		return m.viewError()
	}
	
	return ""
}
