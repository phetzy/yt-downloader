package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
)

// qualityItem implements list.Item for quality selection
type qualityItem struct {
	format FormatInfo
	index  int
}

// FilterValue returns the value to filter on
func (i qualityItem) FilterValue() string {
	return i.format.Quality
}

// Title returns the title of the item
func (i qualityItem) Title() string {
	if i.format.IsAudioOnly {
		return fmt.Sprintf("🎵 %s", i.format.Quality)
	}
	return fmt.Sprintf("📹 %s", i.format.Quality)
}

// Description returns the description of the item
func (i qualityItem) Description() string {
	size := formatBytes(i.format.FileSize)
	
	if i.format.IsAudioOnly {
		return fmt.Sprintf("%s - %s", i.format.Format, size)
	}
	
	if i.format.Resolution != "" {
		return fmt.Sprintf("%s - %s - %s", i.format.Resolution, i.format.Format, size)
	}
	
	return fmt.Sprintf("%s - %s", i.format.Format, size)
}

// convertFormatsToItems converts FormatInfo slice to list items
func convertFormatsToItems(formats []FormatInfo) []list.Item {
	items := make([]list.Item, len(formats))
	for i, format := range formats {
		items[i] = qualityItem{
			format: format,
			index:  i,
		}
	}
	return items
}
