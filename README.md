# RSS Terminal Reader

A simple terminal-based RSS feed reader written in Go. This project uses only the Go standard library and provides a clean terminal interface for reading RSS feeds.

## Features

- Fetch and parse RSS feeds
- Display feed titles and descriptions
- Show latest articles with publication dates
- Clean terminal-based interface

## Installation

1. Ensure you have Go 1.24 or later installed
2. Clone this repository:
```bash
git clone https://github.com/yourusername/rss-terminal.git
cd rss-terminal
```

## Usage

Run the program:
```bash
go run main.go
```

The program will fetch and display the latest articles from Hacker News by default.

## Development

This is a learning project to understand:
- RSS feed parsing with Go's xml package
- HTTP requests in Go
- Terminal UI basics
- Error handling practices

## License

GNU General Public License v3.0