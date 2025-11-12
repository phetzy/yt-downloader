package utils

import (
	"testing"
)

func TestFormatBytes(t *testing.T) {
	tests := []struct {
		name  string
		bytes int64
		want  string
	}{
		{
			name:  "Zero bytes",
			bytes: 0,
			want:  "0 B",
		},
		{
			name:  "Bytes",
			bytes: 512,
			want:  "512 B",
		},
		{
			name:  "Kilobytes",
			bytes: 1536,
			want:  "1.50 KB",
		},
		{
			name:  "Megabytes",
			bytes: 5242880,
			want:  "5.00 MB",
		},
		{
			name:  "Gigabytes",
			bytes: 2147483648,
			want:  "2.00 GB",
		},
		{
			name:  "Terabytes",
			bytes: 1099511627776,
			want:  "1.00 TB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBytes(tt.bytes); got != tt.want {
				t.Errorf("FormatBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatSpeed(t *testing.T) {
	tests := []struct {
		name         string
		bytesPerSec  float64
		want         string
	}{
		{
			name:        "Bytes per second",
			bytesPerSec: 512.5,
			want:        "512 B/s",
		},
		{
			name:        "Kilobytes per second",
			bytesPerSec: 1536.0,
			want:        "1.50 KB/s",
		},
		{
			name:        "Megabytes per second",
			bytesPerSec: 5242880.0,
			want:        "5.00 MB/s",
		},
		{
			name:        "Gigabytes per second",
			bytesPerSec: 2147483648.0,
			want:        "2.00 GB/s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatSpeed(tt.bytesPerSec); got != tt.want {
				t.Errorf("FormatSpeed() = %v, want %v", got, tt.want)
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
			name:    "Zero seconds",
			seconds: 0,
			want:    "0s",
		},
		{
			name:    "Less than a minute",
			seconds: 45,
			want:    "45s",
		},
		{
			name:    "One minute",
			seconds: 60,
			want:    "1m 0s",
		},
		{
			name:    "Several minutes",
			seconds: 195,
			want:    "3m 15s",
		},
		{
			name:    "One hour",
			seconds: 3600,
			want:    "1h 0m 0s",
		},
		{
			name:    "Over one hour",
			seconds: 3665,
			want:    "1h 1m 5s",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDuration(tt.seconds); got != tt.want {
				t.Errorf("FormatDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDurationHMS(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDurationHMS(tt.seconds); got != tt.want {
				t.Errorf("FormatDurationHMS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		maxLen int
		want   string
	}{
		{
			name:   "String shorter than max",
			s:      "Hello",
			maxLen: 10,
			want:   "Hello",
		},
		{
			name:   "String equal to max",
			s:      "HelloWorld",
			maxLen: 10,
			want:   "HelloWorld",
		},
		{
			name:   "String longer than max",
			s:      "Hello World, this is a test",
			maxLen: 15,
			want:   "Hello World,...",
		},
		{
			name:   "Very short max length",
			s:      "Hello",
			maxLen: 3,
			want:   "Hel",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncateString(tt.s, tt.maxLen); got != tt.want {
				t.Errorf("TruncateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
