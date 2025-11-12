# YouTube TUI Downloader - Development Plan

## Project Overview
Build a cross-platform Terminal User Interface (TUI) application in Go that allows users to download YouTube videos with an intuitive, user-friendly interface. The application will leverage Charm's ecosystem for beautiful UI components and animations.

## Technology Stack

### Core Technologies
- **Language**: Go (Golang) 1.21+
- **TUI Framework**: Bubble Tea (event-driven TUI framework)
- **Styling**: Lip Gloss (terminal styling and layout)
- **Components**: Bubbles (pre-built TUI components)
- **YouTube Downloader**: kkdai/youtube (Go library for YouTube downloads)

### Additional Libraries
- **Progress Tracking**: Bubbles progress component
- **File Picker**: Custom implementation with cross-platform support
- **HTTP Client**: Standard library with timeout support
- **Path Handling**: filepath package for cross-platform paths

## Step-by-Step Development Plan

### Phase 1: Project Setup and Architecture (Days 1-2)

#### Step 1.1: Initialize Project Structure
```
yt-downloader/
├── main.go                 # Entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies
├── README.md               # User documentation
├── LICENSE                 # License file
├── .gitignore             # Git ignore rules
├── cmd/
│   └── yt-downloader/
│       └── main.go        # CLI entry point
├── internal/
│   ├── tui/               # TUI components
│   │   ├── app.go         # Main application model
│   │   ├── styles.go      # Lip Gloss styles
│   │   ├── input.go       # URL input screen
│   │   ├── quality.go     # Quality selection screen
│   │   ├── picker.go      # Directory picker screen
│   │   └── download.go    # Download progress screen
│   ├── youtube/           # YouTube operations
│   │   ├── client.go      # YouTube client wrapper
│   │   ├── formats.go     # Format parsing and filtering
│   │   └── downloader.go  # Download logic
│   └── utils/             # Utility functions
│       ├── filepath.go    # Cross-platform path handling
│       └── format.go      # Size/speed formatting
└── assets/                # Optional: screenshots, demo GIFs
```

#### Step 1.2: Initialize Go Module
- Run `go mod init github.com/yourusername/yt-downloader`
- Add core dependencies:
  - `github.com/charmbracelet/bubbletea`
  - `github.com/charmbracelet/bubbles`
  - `github.com/charmbracelet/lipgloss`
  - `github.com/kkdai/youtube/v2`

#### Step 1.3: Create Base Application Structure
- Set up main.go with basic Bubble Tea program
- Define application states (enum):
  - `StateURLInput`: Waiting for YouTube URL
  - `StateLoading`: Fetching video information
  - `StateQualitySelect`: Choosing quality/format
  - `StateDirectoryPicker`: Selecting download location
  - `StateDownloading`: Active download with progress
  - `StateComplete`: Download finished
  - `StateError`: Error occurred

### Phase 2: UI Components Development (Days 3-5)

#### Step 2.1: Design Color Scheme and Styles
Create a cohesive design system using Lip Gloss:
- **Primary Colors**: Define brand colors (e.g., YouTube red #FF0000)
- **Accent Colors**: Complementary colors for highlights
- **Status Colors**: Success (green), Error (red), Info (blue), Warning (yellow)
- **Text Styles**: Title, subtitle, body, help text
- **Border Styles**: Rounded, bold, subtle
- **Layout**: Margins, padding, alignment

#### Step 2.2: Build URL Input Screen
Features:
- Welcome banner with ASCII art or styled title
- Text input field for YouTube URL (using bubbles/textinput)
- Input validation (basic format check)
- Instructions and help text
- Keyboard shortcuts display (Enter to submit, Ctrl+C to quit)

Validation:
- Check for youtube.com or youtu.be domains
- Support various YouTube URL formats:
  - `https://www.youtube.com/watch?v=VIDEO_ID`
  - `https://youtu.be/VIDEO_ID`
  - `https://www.youtube.com/watch?v=VIDEO_ID&list=...`

#### Step 2.3: Create Loading Screen
Features:
- Animated spinner (using bubbles/spinner)
- Status message: "Fetching video information..."
- Ability to cancel operation

#### Step 2.4: Design Quality Selection Screen
Features:
- Display video metadata:
  - Title (truncated if too long)
  - Author/Channel
  - Duration
  - Upload date
  - View count
- List of available formats (using bubbles/list):
  - Video + Audio options (e.g., "1080p MP4", "720p MP4")
  - Audio-only options (e.g., "Audio - High", "Audio - Medium")
- Format details:
  - Resolution
  - Format type (MP4, WebM, etc.)
  - Approximate file size
- Keyboard navigation:
  - Up/Down arrows to select
  - Enter to confirm
  - Esc to go back

#### Step 2.5: Implement Directory Picker
Features:
- Show current directory path
- List directories in current location
- Navigation:
  - Up/Down to select directory
  - Enter to navigate into directory
  - Backspace/Left to go to parent directory
  - Right arrow or Enter on ".." to go up
- Show disk space available
- "Select this directory" option
- Cross-platform path handling
- Default to user's Downloads or Videos folder
- Remember last used directory (optional: use config file)

Special considerations:
- Windows: Handle drive letters (C:\, D:\, etc.)
- macOS/Linux: Start from home directory (~/)
- Show hidden folders option (toggle with 'h')
- Handle permissions gracefully

#### Step 2.6: Build Download Progress Screen
Features:
- Progress bar (using bubbles/progress)
- Download statistics:
  - Percentage complete
  - Download speed (MB/s)
  - Downloaded size / Total size
  - Estimated time remaining
- Animated status indicator
- Ability to cancel download
- Post-download actions:
  - Success message
  - File location display
  - Options: Download another, Open folder, Exit

### Phase 3: YouTube Integration (Days 6-8)

#### Step 3.1: Create YouTube Client Wrapper
Implement client.go:
- Initialize YouTube client
- Fetch video information by URL
- Extract video ID from various URL formats
- Handle authentication if needed (for age-restricted content)
- Error handling for:
  - Invalid URLs
  - Private/unavailable videos
  - Regional restrictions
  - Network issues

#### Step 3.2: Parse and Filter Formats
Implement formats.go:
- Parse available formats from video metadata
- Categorize formats:
  - Video + Audio combined formats
  - Video-only formats (may need merging)
  - Audio-only formats
- Filter by quality:
  - Sort by resolution (descending)
  - Group by codec
  - Prefer formats with audio included
- Calculate or estimate file sizes
- Present user-friendly format descriptions

Format categories:
- **Video + Audio**: 1080p, 720p, 480p, 360p
- **Audio Only**: High (128kbps+), Medium (64-128kbps), Low (<64kbps)

#### Step 3.3: Implement Download Logic
Implement downloader.go:
- Download selected format
- Track progress (bytes downloaded)
- Calculate download speed
- Support resumable downloads (if possible)
- Handle merge operations:
  - If video-only + audio-only formats selected
  - Use ffmpeg for merging (if available)
  - Fallback to combined formats if ffmpeg not found
- Validate downloaded file integrity
- Error handling and retry logic

### Phase 4: State Management and Flow (Days 9-10)

#### Step 4.1: Implement Bubble Tea Model
Create app.go with:
- Main Model struct containing:
  - Current state
  - URL input component
  - Quality list component
  - Directory picker component
  - Progress bar component
  - Video metadata
  - Selected format
  - Download path
  - Error messages
- Init() function: Initial setup
- Update() function: Handle messages/events
- View() function: Render current state

#### Step 4.2: Define Messages (Events)
Custom messages for async operations:
- `urlSubmittedMsg`: URL input submitted
- `videoInfoMsg`: Video info fetched successfully
- `videoInfoErrorMsg`: Failed to fetch video info
- `formatSelectedMsg`: User selected a format
- `directorySelectedMsg`: User selected download location
- `downloadProgressMsg`: Progress update
- `downloadCompleteMsg`: Download finished
- `downloadErrorMsg`: Download failed
- `cancelMsg`: User cancelled operation

#### Step 4.3: Implement State Transitions
State machine logic:
```
StateURLInput → StateLoading → StateQualitySelect → StateDirectoryPicker → StateDownloading → StateComplete
                    ↓                  ↓                       ↓                  ↓               ↓
                StateError         StateError              StateError         StateError    StateURLInput
```

Handle edge cases:
- User cancellation at any stage
- Network timeouts
- Invalid selections
- Disk full errors

### Phase 5: Cross-Platform Compatibility (Day 11)

#### Step 5.1: Path Handling
Implement filepath.go:
- Use `filepath` package for all path operations
- Convert between OS-specific path separators
- Expand ~ to home directory
- Handle Windows drive letters
- Validate write permissions
- Check available disk space

#### Step 5.2: Terminal Compatibility
- Test on various terminal emulators:
  - Windows: CMD, PowerShell, Windows Terminal
  - macOS: Terminal.app, iTerm2
  - Linux: gnome-terminal, konsole, xterm
- Handle terminal size changes
- Ensure color support detection
- Fallback for terminals without true color

#### Step 5.3: Build for Multiple Platforms
Create build scripts or Makefile:
```bash
# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bin/yt-downloader-windows-amd64.exe

# Build for macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o bin/yt-downloader-darwin-amd64

# Build for macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o bin/yt-downloader-darwin-arm64

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o bin/yt-downloader-linux-amd64
```

### Phase 6: Error Handling and Edge Cases (Day 12)

#### Step 6.1: Comprehensive Error Handling
- Network errors (connection timeout, DNS failure)
- Invalid YouTube URLs
- Age-restricted or private videos
- Geo-blocked content
- Insufficient disk space
- Write permission errors
- Interrupted downloads
- Invalid format selection

#### Step 6.2: User-Friendly Error Messages
- Convert technical errors to plain language
- Provide actionable solutions
- Show help text for common issues
- Log detailed errors for debugging (optional)

#### Step 6.3: Graceful Degradation
- If video + audio merge fails, offer separate downloads
- If high quality unavailable, suggest alternatives
- If ffmpeg not found, explain limitation

### Phase 7: Polish and User Experience (Day 13)

#### Step 7.1: Add Help and Documentation
- Built-in help screen (toggle with '?')
- Keyboard shortcuts reference
- Tooltips and inline help
- README with examples

#### Step 7.2: Improve Animations
- Smooth transitions between states
- Loading spinners
- Progress bar animations
- Celebrate download completion (optional: confetti effect)

#### Step 7.3: Configuration (Optional)
Create ~/.yt-downloader/config.yaml:
- Default download directory
- Preferred quality/format
- Theme preferences
- Last used settings

### Phase 8: Testing and Documentation (Days 14-15)

#### Step 8.1: Testing
- Unit tests for core functions:
  - URL parsing
  - Format filtering
  - Size calculations
- Integration tests:
  - Mock YouTube API responses
  - Test state transitions
- Manual testing:
  - Various video types (short, long, live)
  - Different URLs formats
  - All quality options
  - All platforms

#### Step 8.2: Documentation
- Complete README.md
- Usage examples with screenshots/GIFs
- Troubleshooting section
- Contributing guidelines
- Changelog

#### Step 8.3: Release Preparation
- Version tagging (semantic versioning)
- Release notes
- Pre-built binaries for all platforms
- Installation instructions:
  - Direct binary download
  - `go install` command
  - Package managers (brew, scoop, apt - future)

## Key Features Checklist

### Must-Have Features
- ✅ Cross-platform support (Windows, macOS, Linux)
- ✅ YouTube URL input with validation
- ✅ Display video metadata (title, duration, channel)
- ✅ List available qualities/resolutions
- ✅ Video + Audio download option
- ✅ Audio-only download option
- ✅ User-friendly directory picker
- ✅ Real-time download progress with animations
- ✅ Cancel download capability
- ✅ Error handling with clear messages

### Nice-to-Have Features
- ⭐ Resume interrupted downloads
- ⭐ Download queue (multiple videos)
- ⭐ Playlist support
- ⭐ Format conversion (using ffmpeg)
- ⭐ Subtitle download
- ⭐ Configuration file for defaults
- ⭐ Download history
- ⭐ Thumbnail preview (ASCII art or kitty protocol)
- ⭐ Search YouTube from TUI
- ⭐ Copy URL from clipboard automatically

## Development Timeline

- **Week 1**: Setup, UI components, YouTube integration
- **Week 2**: State management, cross-platform testing, polish
- **Week 3**: Documentation, testing, release

Total estimated time: **15 days** (can be compressed or extended based on feature scope)

## Success Criteria

1. ✅ Application runs on Windows, macOS, and Linux without modification
2. ✅ User can paste a YouTube URL and see available formats
3. ✅ User can select video quality and download type
4. ✅ User can choose download location through intuitive interface
5. ✅ Download progress is clearly visible and accurate
6. ✅ Errors are handled gracefully with helpful messages
7. ✅ UI is responsive and visually appealing
8. ✅ README provides clear installation and usage instructions

## Future Enhancements (Post-v1.0)

1. **Multiple Downloads**: Queue system for batch downloads
2. **Playlist Support**: Download entire playlists
3. **Format Conversion**: Built-in audio/video conversion
4. **Themes**: Multiple color themes
5. **Clipboard Integration**: Auto-detect URLs in clipboard
6. **Update Checker**: Notify users of new versions
7. **Proxy Support**: Download through proxy
8. **Advanced Filters**: Filter by codec, bitrate, etc.

## Phase 9: Testing Strategy (Days 16-18)

### Step 9.1: Unit Testing

#### Test Structure
```
yt-downloader/
├── internal/
│   ├── tui/
│   │   ├── app_test.go
│   │   ├── styles_test.go
│   │   └── ...
│   ├── youtube/
│   │   ├── client_test.go
│   │   ├── formats_test.go
│   │   └── downloader_test.go
│   └── utils/
│       ├── filepath_test.go
│       └── format_test.go
└── testdata/              # Test fixtures and mock data
    ├── mock_videos.json
    └── sample_responses/
```

#### Critical Unit Tests

**YouTube Client Tests (`youtube/client_test.go`)**:
- `TestExtractVideoID`: Test video ID extraction from various URL formats
  - Standard format: `https://www.youtube.com/watch?v=VIDEO_ID`
  - Short format: `https://youtu.be/VIDEO_ID`
  - With parameters: `https://www.youtube.com/watch?v=VIDEO_ID&list=...`
  - Edge cases: Invalid URLs, malformed URLs
- `TestFetchVideoInfo`: Test fetching video metadata (using mocked responses)
- `TestHandleErrors`: Test error handling for unavailable/private videos

**Format Parser Tests (`youtube/formats_test.go`)**:
- `TestParseFormats`: Test parsing of format data structures
- `TestFilterVideoFormats`: Test filtering video+audio combined formats
- `TestFilterAudioFormats`: Test filtering audio-only formats
- `TestSortByQuality`: Test quality sorting (descending resolution)
- `TestFormatSizeCalculation`: Test file size estimation
- `TestFormatDescription`: Test user-friendly format descriptions

**Downloader Tests (`youtube/downloader_test.go`)**:
- `TestDownloadProgress`: Test progress tracking calculations
- `TestSpeedCalculation`: Test download speed calculation
- `TestETACalculation`: Test estimated time remaining
- `TestCancelDownload`: Test cancellation logic
- `TestResumeDownload`: Test resume capability (if implemented)

**Utility Tests (`utils/filepath_test.go`)**:
- `TestExpandHomeDir`: Test ~ expansion on all platforms
- `TestNormalizePath`: Test path normalization
- `TestValidateDirectory`: Test directory validation
- `TestCheckDiskSpace`: Test disk space checking
- `TestCrossplatformPaths`: Test Windows/Unix path handling

**Format Utility Tests (`utils/format_test.go`)**:
- `TestFormatBytes`: Test byte size formatting (KB, MB, GB)
- `TestFormatSpeed`: Test speed formatting (MB/s)
- `TestFormatDuration`: Test duration formatting (HH:MM:SS)

#### Testing Approach
- Use **table-driven tests** for multiple test cases
- Mock external dependencies (YouTube API, filesystem)
- Aim for **>80% code coverage** on critical paths
- Use `go test -v ./...` to run all tests
- Use `go test -cover ./...` to check coverage

### Step 9.2: Integration Testing

#### Integration Test Scenarios

**End-to-End Flow Tests**:
- `TestCompleteDownloadFlow`: Full workflow from URL input to download
- `TestStateTransitions`: Test all state machine transitions
- `TestErrorRecovery`: Test recovery from various error states

**Component Integration Tests**:
- `TestUIWithBackend`: Test TUI components with actual YouTube client
- `TestDownloadWithProgress`: Test downloader with progress updates
- `TestFilePickerWithFilesystem`: Test directory picker with real filesystem

#### Mock Data Strategy
- Create test fixtures with sample video metadata
- Mock HTTP responses for YouTube API calls
- Use `httptest` package for HTTP mocking
- Create temporary directories for download tests

### Step 9.3: Manual Testing Checklist

#### Functional Testing
- [ ] URL Input
  - [ ] Valid YouTube URLs work
  - [ ] Invalid URLs show appropriate errors
  - [ ] Clipboard paste works (if implemented)
  - [ ] Various URL formats accepted

- [ ] Video Information Display
  - [ ] Title displays correctly (handles Unicode, long titles)
  - [ ] Duration formatted properly
  - [ ] View count and upload date shown
  - [ ] Author/channel name displayed

- [ ] Quality Selection
  - [ ] All available qualities listed
  - [ ] Audio-only options separated
  - [ ] Keyboard navigation works smoothly
  - [ ] Selection highlights correctly
  - [ ] Back navigation works

- [ ] Directory Picker
  - [ ] Navigation works (up/down, enter, back)
  - [ ] Current path displays correctly
  - [ ] Hidden folders toggle works
  - [ ] Disk space shown accurately
  - [ ] Selection confirmation works

- [ ] Download Progress
  - [ ] Progress bar updates smoothly
  - [ ] Speed calculation accurate
  - [ ] ETA calculation reasonable
  - [ ] Percentage displays correctly
  - [ ] Cancel works at any point
  - [ ] Completion notification shows

- [ ] Error Handling
  - [ ] Network errors handled gracefully
  - [ ] Invalid video errors clear
  - [ ] Permission errors explained
  - [ ] Disk full errors caught
  - [ ] All errors allow retry or exit

#### Cross-Platform Testing

**Windows Testing** (Windows 10/11):
- [ ] CMD prompt
- [ ] PowerShell
- [ ] Windows Terminal
- [ ] Test drive letter handling (C:\, D:\)
- [ ] Test long file paths
- [ ] Test special characters in filenames

**macOS Testing** (macOS 12+):
- [ ] Terminal.app
- [ ] iTerm2
- [ ] Test Intel and Apple Silicon builds
- [ ] Test home directory expansion (~)
- [ ] Test Unicode in paths

**Linux Testing** (Ubuntu 22.04, Fedora 38):
- [ ] gnome-terminal
- [ ] konsole
- [ ] xterm
- [ ] Test various distributions
- [ ] Test different shells (bash, zsh, fish)

#### Performance Testing
- [ ] Large file downloads (>1GB)
- [ ] Multiple quality options (>20 formats)
- [ ] Deep directory navigation (>10 levels)
- [ ] Terminal resize handling
- [ ] Low bandwidth scenarios

#### Accessibility Testing
- [ ] Keyboard-only navigation
- [ ] High contrast themes
- [ ] Screen reader compatibility (limited)
- [ ] Color blind friendly

### Step 9.4: Test Automation Tools

**Testing Libraries**:
- `testing` - Go standard library
- `github.com/stretchr/testify` - Assertions and mocks
- `github.com/golang/mock` - Mock generation (optional)
- `github.com/netflix/go-expect` - TUI testing

**Coverage Tools**:
```bash
# Generate coverage report
go test -coverprofile=coverage.out ./...

# View coverage in browser
go tool cover -html=coverage.out

# Check coverage percentage
go tool cover -func=coverage.out
```

**Benchmarking**:
```bash
# Run benchmarks
go test -bench=. ./...

# With memory profiling
go test -bench=. -benchmem ./...
```

## Phase 10: CI/CD Pipeline Setup (Days 19-20)

### Step 10.1: GitHub Actions Workflow Structure

#### File Structure
```
.github/
├── workflows/
│   ├── ci.yml              # Continuous Integration
│   ├── release.yml         # Release builds
│   └── lint.yml            # Linting and formatting
├── dependabot.yml          # Dependency updates
└── CODEOWNERS             # Code review assignments
```

### Step 10.2: Continuous Integration Pipeline

#### CI Workflow (`.github/workflows/ci.yml`)

**Triggers**:
- Push to `main` and `develop` branches
- Pull requests to `main`
- Manual workflow dispatch

**Jobs**:

1. **Lint and Format Check**
   ```yaml
   - Run golangci-lint
   - Run gofmt check
   - Run go vet
   - Check for security issues (gosec)
   ```

2. **Test on Multiple Platforms**
   - Matrix strategy:
     - OS: Ubuntu (latest), macOS (latest), Windows (latest)
     - Go version: 1.21, 1.22, 1.23
   - Steps:
     - Checkout code
     - Setup Go
     - Cache dependencies
     - Download modules
     - Run tests with race detector
     - Upload coverage to Codecov

3. **Build Verification**
   - Build for all target platforms:
     - linux/amd64, linux/arm64
     - darwin/amd64, darwin/arm64
     - windows/amd64, windows/arm64
   - Ensure binaries are created
   - Check binary sizes

**Example CI Workflow**:
```yaml
name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
        go-version: ['1.21', '1.22']
    runs-on: ${{ matrix.os }}
    
    steps:
    - uses: actions/checkout@v4
    
    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: ${{ matrix.go-version }}
        cache: true
    
    - name: Download dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v -race -coverprofile=coverage.out ./...
    
    - name: Upload coverage
      uses: codecov/codecov-action@v4
      with:
        file: ./coverage.out

  lint:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    
    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.22'
    
    - name: golangci-lint
      uses: golangci/golangci-lint-action@v4
      with:
        version: latest
```

### Step 10.3: Continuous Deployment Pipeline

#### Release Workflow (`.github/workflows/release.yml`)

**Triggers**:
- Push of version tags (v*.*.* pattern)
- Manual workflow dispatch with version input

**Jobs**:

1. **Create Release Binaries**
   - Matrix build for all platforms:
     - `windows/amd64` → `yt-downloader-windows-amd64.exe`
     - `darwin/amd64` → `yt-downloader-darwin-amd64`
     - `darwin/arm64` → `yt-downloader-darwin-arm64`
     - `linux/amd64` → `yt-downloader-linux-amd64`
     - `linux/arm64` → `yt-downloader-linux-arm64`
   - Build flags:
     - Strip debug info: `-ldflags="-s -w"`
     - Add version info: `-ldflags="-X main.version=$VERSION"`
   - Compress binaries (UPX optional)

2. **Generate Checksums**
   - Create SHA256 checksums for all binaries
   - Create `checksums.txt` file

3. **Create GitHub Release**
   - Generate changelog from commits
   - Upload all binaries
   - Upload checksums
   - Mark as pre-release if beta version

**Example Release Workflow**:
```yaml
name: Release

on:
  push:
    tags:
      - 'v*.*.*'

jobs:
  release:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
      with:
        fetch-depth: 0
    
    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.22'
    
    - name: Run GoReleaser
      uses: goreleaser/goreleaser-action@v5
      with:
        version: latest
        args: release --clean
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

### Step 10.4: GoReleaser Configuration

**Create `.goreleaser.yml`**:
```yaml
project_name: yt-downloader

before:
  hooks:
    - go mod tidy
    - go test ./...

builds:
  - id: yt-downloader
    binary: yt-downloader
    env:
      - CGO_ENABLED=0
    goos:
      - linux
      - darwin
      - windows
    goarch:
      - amd64
      - arm64
    ldflags:
      - -s -w
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}
    ignore:
      - goos: windows
        goarch: arm64

archives:
  - id: default
    format: tar.gz
    format_overrides:
      - goos: windows
        format: zip
    name_template: >-
      {{ .ProjectName }}_
      {{- .Version }}_
      {{- title .Os }}_
      {{- if eq .Arch "amd64" }}x86_64
      {{- else if eq .Arch "386" }}i386
      {{- else }}{{ .Arch }}{{ end }}

checksum:
  name_template: 'checksums.txt'

snapshot:
  name_template: "{{ incpatch .Version }}-next"

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
      - '^chore:'
      - 'README'
      - 'typo'

release:
  github:
    owner: phetzy
    name: yt-downloader
  draft: false
  prerelease: auto
```

### Step 10.5: Additional CI/CD Features

#### Linting Workflow (`.github/workflows/lint.yml`)
```yaml
name: Lint

on:
  pull_request:

jobs:
  golangci:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v4
    - uses: actions/setup-go@v5
      with:
        go-version: '1.22'
    - name: golangci-lint
      uses: golangci/golangci-lint-action@v4
```

#### Dependabot Configuration (`.github/dependabot.yml`)
```yaml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 10
  
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
```

### Step 10.6: Quality Gates and Badges

#### Badge Configuration (for README)
```markdown
![Build Status](https://github.com/phetzy/yt-downloader/workflows/CI/badge.svg)
![Coverage](https://codecov.io/gh/phetzy/yt-downloader/branch/main/graph/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/phetzy/yt-downloader)
![Latest Release](https://img.shields.io/github/v/release/phetzy/yt-downloader)
```

#### Quality Requirements
- Tests must pass on all platforms
- Code coverage > 80% for core packages
- No security vulnerabilities (gosec)
- Go Report Card grade: A
- All linters pass

### Step 10.7: Release Process

#### Semantic Versioning
- **v1.0.0**: Major release
- **v1.1.0**: Minor release (new features)
- **v1.0.1**: Patch release (bug fixes)

#### Release Checklist
1. Update CHANGELOG.md
2. Update version in code
3. Create and push version tag:
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```
4. GitHub Actions automatically:
   - Runs all tests
   - Builds binaries for all platforms
   - Creates GitHub release
   - Uploads binaries and checksums
5. Update release notes on GitHub
6. Announce release

#### Pre-release Testing
- Test on all target platforms
- Run full test suite
- Manual smoke testing
- Security scan
- Performance benchmarks

## Notes and Considerations

### Technical Decisions
- **Why kkdai/youtube?**: Actively maintained, no external dependencies, good format support
- **Why Bubble Tea?**: Production-ready, excellent documentation, active community
- **Why not ffmpeg bundling?**: Keep binary size small, user can install separately if needed

### Legal and Ethical Considerations
- Add disclaimer about YouTube Terms of Service
- Respect copyright and content creator rights
- Download for personal use only
- Consider rate limiting to avoid abuse

### Performance Considerations
- Handle large files (multi-GB) efficiently
- Don't load entire file into memory
- Stream downloads to disk
- Efficient terminal rendering (avoid redraws)

### Accessibility
- Ensure keyboard-only navigation
- Clear focus indicators
- High contrast color schemes
- Screen reader considerations (limited in TUI)

---

**This plan is a living document. Adjust priorities and timeline based on progress and user feedback.**
