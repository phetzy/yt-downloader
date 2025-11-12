package youtube

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/kkdai/youtube/v2"
)

// Downloader handles downloading YouTube videos
type Downloader struct {
	client *Client
}

// NewDownloader creates a new Downloader instance
func NewDownloader(client *Client) *Downloader {
	return &Downloader{
		client: client,
	}
}

// DownloadProgress represents download progress information
type DownloadProgress struct {
	BytesDownloaded int64
	TotalBytes      int64
	Percentage      float64
	Speed           float64 // bytes per second
	ETA             int     // seconds remaining
	StartTime       time.Time
}

// ProgressCallback is called periodically during download
type ProgressCallback func(progress DownloadProgress)

// Download downloads a video in the specified format to the given path
func (d *Downloader) Download(ctx context.Context, videoID string, format Format, outputPath string, callback ProgressCallback) error {
	// Get video information
	video, err := d.client.client.GetVideo(videoID)
	if err != nil {
		return fmt.Errorf("failed to get video: %w", err)
	}

	// Find the matching format
	var selectedFormat *youtube.Format
	for _, f := range video.Formats {
		if f.ItagNo == format.ItagNo {
			selectedFormat = &f
			break
		}
	}

	if selectedFormat == nil {
		return fmt.Errorf("format not found")
	}

	// Create output file
	outputFile := filepath.Join(outputPath, sanitizeFilename(video.Title)+"."+format.Extension)
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// Get the stream
	stream, _, err := d.client.client.GetStream(video, selectedFormat)
	if err != nil {
		return fmt.Errorf("failed to get stream: %w", err)
	}
	defer stream.Close()

	// Download with progress tracking
	return d.downloadWithProgress(ctx, stream, file, selectedFormat.ContentLength, callback)
}

// downloadWithProgress downloads from a stream with progress tracking
func (d *Downloader) downloadWithProgress(ctx context.Context, reader io.Reader, writer io.Writer, totalSize int64, callback ProgressCallback) error {
	buffer := make([]byte, 32*1024) // 32KB buffer
	var downloaded int64
	startTime := time.Now()
	lastUpdate := startTime

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		n, err := reader.Read(buffer)
		if n > 0 {
			_, writeErr := writer.Write(buffer[:n])
			if writeErr != nil {
				return writeErr
			}

			downloaded += int64(n)

			// Update progress every 100ms
			now := time.Now()
			if now.Sub(lastUpdate) >= 100*time.Millisecond || downloaded == totalSize {
				elapsed := now.Sub(startTime).Seconds()
				speed := float64(downloaded) / elapsed
				percentage := float64(downloaded) / float64(totalSize) * 100

				var eta int
				if speed > 0 {
					remaining := totalSize - downloaded
					eta = int(float64(remaining) / speed)
				}

				if callback != nil {
					callback(DownloadProgress{
						BytesDownloaded: downloaded,
						TotalBytes:      totalSize,
						Percentage:      percentage,
						Speed:           speed,
						ETA:             eta,
						StartTime:       startTime,
					})
				}

				lastUpdate = now
			}
		}

		if err != nil {
			if err == io.EOF {
				// Download complete
				if callback != nil {
					callback(DownloadProgress{
						BytesDownloaded: downloaded,
						TotalBytes:      totalSize,
						Percentage:      100,
						Speed:           0,
						ETA:             0,
						StartTime:       startTime,
					})
				}
				return nil
			}
			return err
		}
	}
}

// sanitizeFilename removes invalid characters from filename
func sanitizeFilename(filename string) string {
	// Replace invalid characters with underscore
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := filename

	for _, char := range invalid {
		result = filepath.Base(result)
		for i := 0; i < len(result); i++ {
			if string(result[i]) == char {
				result = result[:i] + "_" + result[i+1:]
			}
		}
	}

	// Limit length to 200 characters
	if len(result) > 200 {
		result = result[:200]
	}

	return result
}
