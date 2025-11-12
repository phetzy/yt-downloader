package youtube

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kkdai/youtube/v2"
)

// Client wraps the YouTube client
type Client struct {
	client youtube.Client
}

// NewClient creates a new YouTube client
func NewClient() *Client {
	return &Client{
		client: youtube.Client{},
	}
}

// VideoInfo contains information about a YouTube video
type VideoInfo struct {
	ID          string
	Title       string
	Author      string
	Duration    string
	Views       uint64
	UploadDate  string
	Description string
	Formats     []Format
}

// Format represents a video/audio format
type Format struct {
	ItagNo      int
	Quality     string
	Resolution  string
	MimeType    string
	Bitrate     int
	FileSize    int64
	IsAudioOnly bool
	HasVideo    bool
	HasAudio    bool
	Extension   string
}

// GetVideoInfo fetches information about a YouTube video
func (c *Client) GetVideoInfo(url string) (*VideoInfo, error) {
	// Extract video ID from URL
	videoID, err := extractVideoID(url)
	if err != nil {
		return nil, err
	}

	// Fetch video information
	video, err := c.client.GetVideo(videoID)
	if err != nil {
		// Check for common error patterns
		errMsg := err.Error()
		
		// Detect age-restricted or sign-in required videos
		if strings.Contains(errMsg, "400") || strings.Contains(errMsg, "403") {
			return nil, fmt.Errorf("video is age-restricted or requires sign-in. This tool cannot download videos that require YouTube authentication. Please try a different video")
		}
		
		// Detect unavailable videos
		if strings.Contains(errMsg, "404") {
			return nil, fmt.Errorf("video not found. It may be private, deleted, or unavailable in your region")
		}
		
		// Generic error with helpful context
		return nil, fmt.Errorf("failed to fetch video info: %w. This may be due to regional restrictions, age restrictions, or the video requiring sign-in", err)
	}

	// Parse formats
	formats := c.parseFormats(video.Formats)

	// Format duration
	duration := formatDuration(int(video.Duration.Seconds()))

	return &VideoInfo{
		ID:          video.ID,
		Title:       video.Title,
		Author:      video.Author,
		Duration:    duration,
		Views:       uint64(video.Views),
		UploadDate:  video.PublishDate.Format("2006-01-02"),
		Description: video.Description,
		Formats:     formats,
	}, nil
}

// parseFormats converts youtube.Format to our Format type
// Prioritizes MP4/M4A formats for better compatibility
func (c *Client) parseFormats(ytFormats youtube.FormatList) []Format {
	var formats []Format

	for _, f := range ytFormats {
		format := Format{
			ItagNo:      f.ItagNo,
			Quality:     f.Quality,
			MimeType:    f.MimeType,
			Bitrate:     f.Bitrate,
			FileSize:    f.ContentLength,
			IsAudioOnly: strings.Contains(f.MimeType, "audio"),
			HasVideo:    strings.Contains(f.MimeType, "video"),
			HasAudio:    f.AudioChannels > 0,
		}

		// Determine resolution and generate proper quality label
		if f.Width > 0 && f.Height > 0 {
			format.Resolution = fmt.Sprintf("%dx%d", f.Width, f.Height)
			// Generate proper quality label from height (720p, 1080p, etc)
			format.Quality = fmt.Sprintf("%dp", f.Height)
		} else if format.IsAudioOnly {
			// For audio-only, use bitrate-based quality
			if f.Bitrate >= 128000 {
				format.Quality = "Audio - High"
			} else if f.Bitrate >= 96000 {
				format.Quality = "Audio - Medium"
			} else {
				format.Quality = "Audio - Low"
			}
		}

		// Determine extension from MIME type
		format.Extension = getExtensionFromMimeType(f.MimeType)

		// Include video formats (combined or video-only) and audio-only formats
		// Modern YouTube often separates video and audio streams
		if format.HasVideo || format.IsAudioOnly {
			// Only include if it has reasonable file size (skip empty streams)
			if format.FileSize > 0 {
				formats = append(formats, format)
			}
		}
	}

	// Sort formats to prioritize MP4/M4A over WebM for compatibility
	// MP4/M4A first, then WebM, then others
	sortFormatsByCompatibility(formats)

	return formats
}

// sortFormatsByCompatibility sorts formats to prioritize more compatible formats
func sortFormatsByCompatibility(formats []Format) {
	// Sort in place - MP4/M4A formats first, then WebM, then others
	for i := 0; i < len(formats); i++ {
		for j := i + 1; j < len(formats); j++ {
			// Prefer MP4/M4A over WebM
			iPreferred := strings.Contains(formats[i].MimeType, "mp4") || 
				strings.Contains(formats[i].MimeType, "m4a")
			jPreferred := strings.Contains(formats[j].MimeType, "mp4") || 
				strings.Contains(formats[j].MimeType, "m4a")
			
			if !iPreferred && jPreferred {
				// Swap - j is more preferred than i
				formats[i], formats[j] = formats[j], formats[i]
			}
		}
	}
}

// extractVideoID extracts the video ID from various YouTube URL formats
func extractVideoID(url string) (string, error) {
	url = strings.TrimSpace(url)

	// Handle different URL formats
	if strings.Contains(url, "youtube.com/watch?v=") {
		parts := strings.Split(url, "v=")
		if len(parts) < 2 {
			return "", errors.New("invalid YouTube URL")
		}
		videoID := strings.Split(parts[1], "&")[0]
		return videoID, nil
	}

	if strings.Contains(url, "youtu.be/") {
		parts := strings.Split(url, "youtu.be/")
		if len(parts) < 2 {
			return "", errors.New("invalid YouTube URL")
		}
		videoID := strings.Split(parts[1], "?")[0]
		return videoID, nil
	}

	if strings.Contains(url, "youtube.com/v/") {
		parts := strings.Split(url, "/v/")
		if len(parts) < 2 {
			return "", errors.New("invalid YouTube URL")
		}
		videoID := strings.Split(parts[1], "?")[0]
		return videoID, nil
	}

	if strings.Contains(url, "youtube.com/embed/") {
		parts := strings.Split(url, "/embed/")
		if len(parts) < 2 {
			return "", errors.New("invalid YouTube URL")
		}
		videoID := strings.Split(parts[1], "?")[0]
		return videoID, nil
	}

	// If none of the patterns match, assume it's just the video ID
	if len(url) == 11 {
		return url, nil
	}

	return "", errors.New("invalid YouTube URL format")
}

// getExtensionFromMimeType extracts file extension from MIME type
func getExtensionFromMimeType(mimeType string) string {
	mimeType = strings.ToLower(mimeType)

	if strings.Contains(mimeType, "mp4") {
		return "mp4"
	}
	if strings.Contains(mimeType, "webm") {
		return "webm"
	}
	if strings.Contains(mimeType, "3gpp") {
		return "3gp"
	}
	if strings.Contains(mimeType, "m4a") {
		return "m4a"
	}
	if strings.Contains(mimeType, "audio") {
		return "m4a"
	}

	return "mp4" // default
}

// formatDuration formats duration in seconds to HH:MM:SS or MM:SS
func formatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutes, secs)
	}
	return fmt.Sprintf("%d:%02d", minutes, secs)
}
