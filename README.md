# 🎬 YouTube TUI Downloader

A beautiful, cross-platform Terminal User Interface (TUI) application for downloading YouTube videos with style and ease.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8.svg)
![Platform](https://img.shields.io/badge/platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey.svg)

> **Built with ❤️ using [Charm](https://charm.sh) tools**

## ✨ Features

- 🎨 **Beautiful TUI** - Intuitive interface built with Bubble Tea, Bubbles, and Lip Gloss
- 🖥️ **Cross-Platform** - Works seamlessly on Windows, macOS, and Linux
- 📊 **Quality Selection** - Choose from all available video qualities and formats
- 🎵 **Audio Downloads** - Extract audio-only from videos
- 📁 **Smart File Picker** - Navigate your filesystem with ease to choose download location
- 📈 **Real-Time Progress** - Watch your download progress with animated progress bars
- ⚡ **Fast & Efficient** - Built in Go for optimal performance
- 🎯 **User-Friendly** - No technical knowledge required - just paste and download!

## 🚀 Quick Start

### Installation

#### Option 1: Download Pre-built Binary (Recommended)

1. Download the latest release for your platform from the [Releases](https://github.com/phetzy/yt-downloader/releases) page:
   - **Windows**: `yt-downloader-windows-amd64.exe`
   - **macOS (Intel)**: `yt-downloader-darwin-amd64`
   - **macOS (Apple Silicon)**: `yt-downloader-darwin-arm64`
   - **Linux**: `yt-downloader-linux-amd64`

2. Make it executable (macOS/Linux only):
   ```bash
   chmod +x yt-downloader-*
   ```

3. Run it:
   ```bash
   ./yt-downloader
   ```

#### Option 2: Install with Go

If you have Go installed (version 1.21 or higher):

```bash
go install github.com/phetzy/yt-downloader@latest
```

#### Option 3: Build from Source

```bash
# Clone the repository
git clone https://github.com/phetzy/yt-downloader.git
cd yt-downloader

# Build the application
go build -o yt-downloader

# Run it
./yt-downloader
```

### Usage

1. **Launch the application**
   ```bash
   yt-downloader
   ```

2. **Paste a YouTube URL**
   - Enter any valid YouTube video URL
   - Formats supported:
     - `https://www.youtube.com/watch?v=VIDEO_ID`
     - `https://youtu.be/VIDEO_ID`
     - URLs with playlists or timestamps

3. **Select Quality**
   - Browse available video qualities (1080p, 720p, 480p, etc.)
   - Choose between:
     - **Video + Audio**: Complete video file
     - **Audio Only**: Extract just the audio

4. **Choose Download Location**
   - Navigate through your directories using arrow keys
   - Press Enter to select the current directory
   - Default location: Your Downloads or Videos folder

5. **Watch the Magic** ✨
   - See real-time download progress
   - View download speed and estimated time
   - Get notified when complete!

## 📸 Screenshots

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│     🎬 YouTube TUI Downloader                              │
│                                                             │
│     Paste your YouTube URL below:                          │
│     ┌─────────────────────────────────────────────────┐   │
│     │ https://www.youtube.com/watch?v=...             │   │
│     └─────────────────────────────────────────────────┘   │
│                                                             │
│     Press Enter to continue • Ctrl+C to quit               │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```

## 🎮 Keyboard Shortcuts

### Global
- `Ctrl+C` or `q` - Quit application
- `?` - Show help screen

### URL Input Screen
- `Enter` - Submit URL
- `Ctrl+U` - Clear input

### Quality Selection Screen
- `↑/↓` or `j/k` - Navigate list
- `Enter` - Select format
- `Esc` - Go back

### Directory Picker Screen
- `↑/↓` or `j/k` - Navigate directories
- `Enter` - Enter directory or select
- `Backspace` or `←` - Go to parent directory
- `h` - Toggle hidden folders
- `Space` - Select current directory

### Download Screen
- `Ctrl+C` - Cancel download
- `Enter` - Download another (when complete)

## 🛠️ Technical Details

### Built With

- **[Go](https://golang.org/)** - Fast, compiled, and cross-platform
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Powerful TUI framework
- **[Bubbles](https://github.com/charmbracelet/bubbles)** - TUI components (inputs, lists, progress bars)
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)** - Style definitions for terminal output
- **[kkdai/youtube](https://github.com/kkdai/youtube)** - YouTube video download library

### Architecture

```
yt-downloader/
├── internal/
│   ├── tui/          # TUI screens and components
│   ├── youtube/      # YouTube client and downloader
│   └── utils/        # Helper functions
└── main.go           # Application entry point
```

### Supported Platforms

| Platform | Architectures | Status |
|----------|---------------|--------|
| Windows  | amd64, arm64  | ✅ Supported |
| macOS    | amd64 (Intel), arm64 (Apple Silicon) | ✅ Supported |
| Linux    | amd64, arm64, 386 | ✅ Supported |

### Video Formats Supported

- **Video Qualities**: 2160p (4K), 1440p, 1080p, 720p, 480p, 360p, 240p, 144p
- **Containers**: MP4, WebM
- **Codecs**: H.264, VP9, AV1
- **Audio**: AAC, Opus, Vorbis

## 📋 Requirements

### Minimum Requirements
- **Terminal**: Any modern terminal emulator with UTF-8 support
- **Colors**: 256-color support recommended (true color supported)
- **Disk Space**: Varies based on video size
- **Internet**: Active internet connection

### Optional Requirements
- **FFmpeg**: Required for merging video-only and audio-only streams (for highest quality downloads)
  - Install on macOS: `brew install ffmpeg`
  - Install on Linux: `sudo apt install ffmpeg` or `sudo yum install ffmpeg`
  - Install on Windows: Download from [ffmpeg.org](https://ffmpeg.org/download.html)

## ❓ FAQ

### Why do I need FFmpeg?

Some high-quality video formats on YouTube separate video and audio streams. FFmpeg is used to merge them into a single file. If FFmpeg is not installed, the application will automatically select combined formats (which may have lower quality).

### Is this legal?

This tool is provided for personal use only. Please respect YouTube's Terms of Service and copyright laws. Only download content you have the right to download, and respect content creators' rights.

### Does this work with private or age-restricted videos?

Currently, the application supports public videos. Support for age-restricted content may be added in future versions.

### Can I download playlists?

Playlist support is planned for a future release. For now, you can download videos one at a time.

### What about subtitles?

Subtitle download support is planned for a future release.

### The download is slow. Why?

Download speed depends on:
- Your internet connection speed
- YouTube's server response time
- Current network congestion

The application downloads as fast as your connection allows.

## 🐛 Troubleshooting

### "Video not available" error
- Check if the URL is correct
- Verify the video is public and not region-locked
- Try again in a few moments (temporary YouTube issue)

### "Permission denied" when saving file
- Ensure you have write permissions to the selected directory
- Try selecting a different download location

### Application crashes on startup
- Ensure your terminal supports UTF-8 and ANSI escape codes
- Try updating to the latest version
- Check if your terminal emulator is supported

### Progress bar doesn't update
- This can happen in some terminal emulators with limited ANSI support
- Try using a modern terminal (Windows Terminal, iTerm2, etc.)

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. **Report Bugs**: Open an issue with details about the problem
2. **Suggest Features**: Share your ideas for improvements
3. **Submit Pull Requests**: Fix bugs or add features
4. **Improve Documentation**: Help make the docs better

### Development Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/yt-downloader.git
cd yt-downloader

# Install dependencies
go mod download

# Run tests
go test ./...

# Build
go build

# Run
./yt-downloader
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Charm](https://charm.sh) for the amazing TUI libraries
- [kkdai](https://github.com/kkdai) for the YouTube download library
- All contributors and users of this project

## 📧 Support

- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/phetzy/yt-downloader/issues)
- 💡 **Feature Requests**: [GitHub Discussions](https://github.com/phetzy/yt-downloader/discussions)
- 📖 **Documentation**: [Wiki](https://github.com/phetzy/yt-downloader/wiki)

## 🗺️ Roadmap

### v1.0 (Current)
- [x] Basic video download
- [x] Quality selection
- [x] Audio-only extraction
- [x] Cross-platform support
- [x] Progress tracking

### v1.1 (Planned)
- [ ] Playlist support
- [ ] Download queue
- [ ] Resume interrupted downloads
- [ ] Subtitle download
- [ ] Configuration file

### v2.0 (Future)
- [ ] Multiple simultaneous downloads
- [ ] Built-in video player
- [ ] Thumbnail preview
- [ ] Search YouTube from TUI
- [ ] Download history

---

<p align="center">
  Made with ❤️ and ☕
  <br>
  <sub>If you find this useful, consider ⭐ starring the repo!</sub>
</p>
