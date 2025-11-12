package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestExpandHomeDir(t *testing.T) {
	homeDir, _ := os.UserHomeDir()
	
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Path with tilde",
			path:    "~/Documents",
			wantErr: false,
		},
		{
			name:    "Just tilde",
			path:    "~",
			wantErr: false,
		},
		{
			name:    "Path without tilde",
			path:    "/usr/local",
			wantErr: false,
		},
		{
			name:    "Relative path",
			path:    "./test",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExpandHomeDir(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExpandHomeDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			// If path starts with ~, result should start with home directory
			if tt.path == "~" && got != homeDir {
				t.Errorf("ExpandHomeDir() = %v, want %v", got, homeDir)
			}
			
			// If path doesn't start with ~, should be unchanged
			if tt.path[0] != '~' && got != tt.path {
				t.Errorf("ExpandHomeDir() = %v, want %v", got, tt.path)
			}
		})
	}
}

func TestNormalizePath(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Path with double slashes",
			path: "/usr//local//bin",
			want: filepath.Clean("/usr/local/bin"),
		},
		{
			name: "Path with dot",
			path: "/usr/./local",
			want: filepath.Clean("/usr/local"),
		},
		{
			name: "Path with double dot",
			path: "/usr/local/../bin",
			want: filepath.Clean("/usr/bin"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizePath(tt.path); got != tt.want {
				t.Errorf("NormalizePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultDownloadDir(t *testing.T) {
	dir, err := GetDefaultDownloadDir()
	if err != nil {
		t.Fatalf("GetDefaultDownloadDir() error = %v", err)
	}

	// Should return a non-empty string
	if dir == "" {
		t.Error("GetDefaultDownloadDir() returned empty string")
	}

	// Should contain "Downloads" on most systems
	if runtime.GOOS != "linux" || os.Getenv("XDG_DOWNLOAD_DIR") == "" {
		// On Windows and macOS, and Linux without XDG, should end with Downloads
		if filepath.Base(dir) != "Downloads" {
			t.Logf("Note: Default download dir is %s (expected to end with 'Downloads')", dir)
		}
	}
}

func TestJoinPath(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		want     string
	}{
		{
			name:     "Simple join",
			elements: []string{"usr", "local", "bin"},
			want:     filepath.Join("usr", "local", "bin"),
		},
		{
			name:     "Join with absolute path",
			elements: []string{"/usr", "local", "bin"},
			want:     filepath.Join("/usr", "local", "bin"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinPath(tt.elements...); got != tt.want {
				t.Errorf("JoinPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHidden(t *testing.T) {
	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "Hidden file with dot",
			path: ".hidden",
			want: true,
		},
		{
			name: "Normal file",
			path: "visible.txt",
			want: false,
		},
		{
			name: "Hidden directory",
			path: ".config",
			want: true,
		},
		{
			name: "Current directory",
			path: ".",
			want: true,
		},
		{
			name: "Parent directory",
			path: "..",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHidden(tt.path); got != tt.want {
				t.Errorf("IsHidden() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetParentDir(t *testing.T) {
	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "Simple path",
			path: "/usr/local/bin",
			want: "/usr/local",
		},
		{
			name: "Root parent",
			path: "/usr",
			want: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetParentDir(tt.path)
			// Normalize both for comparison
			gotNorm := filepath.Clean(got)
			wantNorm := filepath.Clean(tt.want)
			if gotNorm != wantNorm {
				t.Errorf("GetParentDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
