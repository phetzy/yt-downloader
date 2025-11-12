package youtube

import (
	"testing"
)

func TestExtractVideoID(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		want    string
		wantErr bool
	}{
		{
			name:    "Standard YouTube URL",
			url:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
			want:    "dQw4w9WgXcQ",
			wantErr: false,
		},
		{
			name:    "Short YouTube URL",
			url:     "https://youtu.be/dQw4w9WgXcQ",
			want:    "dQw4w9WgXcQ",
			wantErr: false,
		},
		{
			name:    "YouTube URL with parameters",
			url:     "https://www.youtube.com/watch?v=dQw4w9WgXcQ&list=PLrAXtmErZgOeiKm4sgNOknGvNjby9efdf",
			want:    "dQw4w9WgXcQ",
			wantErr: false,
		},
		{
			name:    "YouTube embed URL",
			url:     "https://www.youtube.com/embed/dQw4w9WgXcQ",
			want:    "dQw4w9WgXcQ",
			wantErr: false,
		},
		{
			name:    "YouTube v URL",
			url:     "https://www.youtube.com/v/dQw4w9WgXcQ",
			want:    "dQw4w9WgXcQ",
			wantErr: false,
		},
		{
			name:    "Just video ID",
			url:     "dQw4w9WgXcQ",
			want:    "dQw4w9WgXcQ",
			wantErr: false,
		},
		{
			name:    "Invalid URL",
			url:     "https://example.com/watch?v=test",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Empty URL",
			url:     "",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractVideoID(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractVideoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("extractVideoID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetExtensionFromMimeType(t *testing.T) {
	tests := []struct {
		name     string
		mimeType string
		want     string
	}{
		{
			name:     "MP4 video",
			mimeType: "video/mp4; codecs=\"avc1.42001E, mp4a.40.2\"",
			want:     "mp4",
		},
		{
			name:     "WebM video",
			mimeType: "video/webm; codecs=\"vp9\"",
			want:     "webm",
		},
		{
			name:     "3GP video",
			mimeType: "video/3gpp",
			want:     "3gp",
		},
		{
			name:     "M4A audio",
			mimeType: "audio/mp4; codecs=\"mp4a.40.2\"",
			want:     "m4a",
		},
		{
			name:     "Generic audio",
			mimeType: "audio/webm",
			want:     "m4a",
		},
		{
			name:     "Unknown type",
			mimeType: "application/octet-stream",
			want:     "mp4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getExtensionFromMimeType(tt.mimeType); got != tt.want {
				t.Errorf("getExtensionFromMimeType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name    string
		seconds int
		want    string
	}{
		{
			name:    "Less than a minute",
			seconds: 45,
			want:    "0:45",
		},
		{
			name:    "Exactly one minute",
			seconds: 60,
			want:    "1:00",
		},
		{
			name:    "Several minutes",
			seconds: 195,
			want:    "3:15",
		},
		{
			name:    "One hour",
			seconds: 3600,
			want:    "1:00:00",
		},
		{
			name:    "Over one hour",
			seconds: 3665,
			want:    "1:01:05",
		},
		{
			name:    "Multiple hours",
			seconds: 7384,
			want:    "2:03:04",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatDuration(tt.seconds); got != tt.want {
				t.Errorf("formatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
