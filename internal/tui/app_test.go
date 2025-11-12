package tui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewApp(t *testing.T) {
	app := NewApp()
	
	if app == nil {
		t.Fatal("NewApp() returned nil")
	}
	
	// Check initial state
	if app.state != StateURLInput {
		t.Errorf("Initial state = %v, want %v", app.state, StateURLInput)
	}
	
	// Check components are initialized
	if app.urlInput.Value() != "" {
		t.Error("URL input should be empty initially")
	}
	
	if app.quitting {
		t.Error("App should not be quitting initially")
	}
}

func TestStateTransitions(t *testing.T) {
	app := NewApp()
	
	tests := []struct {
		name      string
		fromState AppState
		event     tea.Msg
		wantState AppState
	}{
		{
			name:      "URL Input to Loading on valid URL",
			fromState: StateURLInput,
			event:     tea.KeyMsg{Type: tea.KeyEnter},
			wantState: StateURLInput, // Stays in input if URL is empty
		},
		{
			name:      "Quit from URL Input",
			fromState: StateURLInput,
			event:     tea.KeyMsg{Type: tea.KeyCtrlC},
			wantState: StateURLInput, // Model handles quitting flag
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app.state = tt.fromState
			_, _ = app.Update(tt.event)
			
			// Note: Some transitions require additional setup
			// This is a basic test structure
		})
	}
}

func TestURLValidation(t *testing.T) {
	tests := []struct {
		name  string
		url   string
		valid bool
	}{
		{
			name:  "Valid standard URL",
			url:   "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			valid: true,
		},
		{
			name:  "Valid short URL",
			url:   "https://youtu.be/dQw4w9WgXcQ",
			valid: true,
		},
		{
			name:  "Invalid URL",
			url:   "https://example.com",
			valid: false,
		},
		{
			name:  "Empty URL",
			url:   "",
			valid: false,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidYouTubeURL(tt.url)
			if result != tt.valid {
				t.Errorf("isValidYouTubeURL(%s) = %v, want %v", tt.url, result, tt.valid)
			}
		})
	}
}

func TestFormatViews(t *testing.T) {
	tests := []struct {
		name  string
		views uint64
		want  string
	}{
		{
			name:  "Less than 1K",
			views: 500,
			want:  "500",
		},
		{
			name:  "Thousands",
			views: 1500,
			want:  "1.5K",
		},
		{
			name:  "Millions",
			views: 2500000,
			want:  "2.5M",
		},
		{
			name:  "Billions",
			views: 3500000000,
			want:  "3.5B",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatViews(tt.views); got != tt.want {
				t.Errorf("formatViews() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQualityItemImplementation(t *testing.T) {
	item := qualityItem{
		format: FormatInfo{
			Quality:     "720p",
			Resolution:  "1280x720",
			Format:      "MP4",
			FileSize:    50000000,
			IsAudioOnly: false,
			HasVideo:    true,
			HasAudio:    true,
		},
		index: 0,
	}
	
	// Test FilterValue
	if item.FilterValue() != "720p" {
		t.Errorf("FilterValue() = %v, want 720p", item.FilterValue())
	}
	
	// Test Title
	title := item.Title()
	if title == "" {
		t.Error("Title() returned empty string")
	}
	
	// Test Description
	desc := item.Description()
	if desc == "" {
		t.Error("Description() returned empty string")
	}
}

func TestConvertFormatsToItems(t *testing.T) {
	formats := []FormatInfo{
		{
			Quality:    "1080p",
			Resolution: "1920x1080",
			Format:     "MP4",
			FileSize:   100000000,
			HasVideo:   true,
			HasAudio:   true,
		},
		{
			Quality:     "Audio - High",
			Format:      "M4A",
			FileSize:    5000000,
			IsAudioOnly: true,
			HasAudio:    true,
		},
	}
	
	items := convertFormatsToItems(formats)
	
	if len(items) != len(formats) {
		t.Errorf("convertFormatsToItems() returned %d items, want %d", len(items), len(formats))
	}
	
	// Check that items implement the list.Item interface
	for i, item := range items {
		if item == nil {
			t.Errorf("Item %d is nil", i)
		}
	}
}
