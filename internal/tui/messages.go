package tui

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

// Custom errors
var (
	ErrInvalidURL      = errors.New("invalid YouTube URL")
	ErrVideoNotFound   = errors.New("video not found or unavailable")
	ErrNetworkError    = errors.New("network error occurred")
	ErrPermissionDenied = errors.New("permission denied")
	ErrDiskFull        = errors.New("not enough disk space")
)

// Custom messages for Bubble Tea

// videoInfoMsg contains the fetched video information
type videoInfoMsg struct {
	Title       string
	Author      string
	Duration    string
	Views       string
	UploadDate  string
	Formats     []FormatInfo
}

// FormatInfo contains information about a video format
type FormatInfo struct {
	ID          string
	Quality     string
	Resolution  string
	Format      string
	FileSize    int64
	IsAudioOnly bool
	HasVideo    bool
	HasAudio    bool
}

// errMsg wraps an error for Bubble Tea
type errMsg struct {
	err error
}

// downloadProgressMsg contains download progress information
type downloadProgressMsg struct {
	BytesDownloaded int64
	TotalBytes      int64
	Speed           float64
	ETA             int
}

// downloadCompleteMsg indicates download completion
type downloadCompleteMsg struct {
	FilePath string
}

// fetchVideoInfo fetches video information from YouTube
// This will be implemented in the youtube package
func fetchVideoInfo(url string) tea.Cmd {
	return func() tea.Msg {
		// TODO: Implement actual YouTube API call
		// For now, return a placeholder
		return videoInfoMsg{
			Title:      "Sample Video",
			Author:     "Sample Channel",
			Duration:   "10:30",
			Views:      "1,000,000",
			UploadDate: "2024-01-01",
			Formats: []FormatInfo{
				{
					ID:         "22",
					Quality:    "720p",
					Resolution: "1280x720",
					Format:     "MP4",
					FileSize:   50000000,
					HasVideo:   true,
					HasAudio:   true,
				},
				{
					ID:         "18",
					Quality:    "360p",
					Resolution: "640x360",
					Format:     "MP4",
					FileSize:   20000000,
					HasVideo:   true,
					HasAudio:   true,
				},
				{
					ID:          "140",
					Quality:     "Audio - High",
					Format:      "M4A",
					FileSize:    5000000,
					IsAudioOnly: true,
					HasAudio:    true,
				},
			},
		}
	}
}
