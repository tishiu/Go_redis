# Redis-like In-Memory Key-Value Store in Go

A lightweight, high-performance in-memory key-value store inspired by Redis, implemented in Go with I/O multiplexing for handling multiple client connections.

## Features

- In-memory key-value storage
- I/O multiplexing for handling multiple clients
- Thread pool for concurrent request handling
- Simple TCP-based protocol

## Prerequisites

- Go 1.16 or higher
- Linux/Unix-like system (for syscall support)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/tishiu/redis_go.git
   cd redis_go
   ```

2. Build the project:
   ```bash
   go build -o redis-go ./cmd
   ```

## Usage

1. Start the server:
   ```bash
   ./redis-go
   ```
   The server will start on `localhost:3000` by default.

2. Connect to the server using `netcat` or `telnet`:
   ```bash
   nc localhost 3000
   ```

## Configuration

You can modify the server configuration in `internal/config/config.go`:
- `Protocol`: Network protocol (default: "tcp")
- `Port`: Server port (default: ":3000")
- `MaxConnection`: Maximum number of concurrent connections (default: 20000)

## Project Structure

- `cmd/`: Main application entry point
- `internal/`: Core application code
  - `server/`: Server implementation
  - `config/`: Configuration settings
  - `core/`: Core functionality
- `TcpServer/`: TCP server implementation
- `ThreadPool/`: Thread pool implementation

## License

MIT
