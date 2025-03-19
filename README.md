# ClipCode

A lightweight clipboard manager written in Go that maintains a history of your clipboard entries and allows quick access through hotkeys.

## Features

- Maintains a history of up to 100 recent clipboard entries
- Quick access to historical entries using Alt+Shift+1 through Alt+Shift+9
- Real-time clipboard monitoring
- Graceful shutdown support

## Prerequisites

- Go 1.16 or higher
- Windows

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/ClipCode.git
cd ClipCode
```

2. Install dependencies:
```bash
go mod download
```

## Running the Application

1. Start the application:
```bash
go run main.go
```

2. The clipboard manager will start running in the background and monitor your clipboard.

## Usage

- Copy any text as you normally would
- Access previous clipboard entries using the following hotkeys:
  - Alt+Shift+1: Paste the most recent clipboard entry
  - Alt+Shift+2: Paste the second most recent entry
  - And so on up to Alt+Shift+9

## Exit

- Press Ctrl+C in the terminal to gracefully shut down the application

## Dependencies

- github.com/atotto/clipboard: For clipboard operations
- github.com/robotn/gohook: For global hotkey support