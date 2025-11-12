package tui

import (
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/phetzy/yt-downloader/internal/youtube"
)

// Custom errors
var (
	ErrInvalidURL           = errors.New("invalid YouTube URL")
	ErrVideoNotFound        = errors.New("video not found or unavailable")
	ErrNetworkError         = errors.New("network error occurred")
	ErrPermissionDenied     = errors.New("permission denied")
	ErrDiskFull             = errors.New("not enough disk space")
	ErrAgeRestricted        = errors.New("video is age-restricted or requires sign-in")
	ErrVideoUnavailable     = errors.New("video is unavailable in your region or requires authentication")
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

// getYouTubeClient creates a new YouTube client instance
func getYouTubeClient() *youtube.Client {
	return youtube.NewClient()
}

// fetchVideoInfo fetches video information from YouTube
func fetchVideoInfo(url string) tea.Cmd {
	return func() tea.Msg {
		// Import the youtube package at the top level
		// This function will be called asynchronously
		client := getYouTubeClient()
		
		videoInfo, err := client.GetVideoInfo(url)
		if err != nil {
			return errMsg{err: err}
		}
		
		// Convert YouTube formats to TUI format
		formats := make([]FormatInfo, len(videoInfo.Formats))
		for i, f := range videoInfo.Formats {
			formats[i] = FormatInfo{
				ID:          f.Quality,
				Quality:     f.Quality,
				Resolution:  f.Resolution,
				Format:      f.Extension,
				FileSize:    f.FileSize,
				IsAudioOnly: f.IsAudioOnly,
				HasVideo:    f.HasVideo,
				HasAudio:    f.HasAudio,
			}
		}
		
		return videoInfoMsg{
			Title:      videoInfo.Title,
			Author:     videoInfo.Author,
			Duration:   videoInfo.Duration,
			Views:      formatViews(videoInfo.Views),
			UploadDate: videoInfo.UploadDate,
			Formats:    formats,
		}
	}
}

// Helper function to format view count
func formatViews(views uint64) string {
	if views >= 1000000000 {
		return fmt.Sprintf("%.1fB", float64(views)/1000000000)
	}
	if views >= 1000000 {
		return fmt.Sprintf("%.1fM", float64(views)/1000000)
	}
	if views >= 1000 {
		return fmt.Sprintf("%.1fK", float64(views)/1000)
	}
	return fmt.Sprintf("%d", views)
}
