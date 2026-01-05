# Go TCP Server

A simple TCP server written in Go that listens on port **9000** and prints messages received from connected clients.

## 📌 Features

- Listens for TCP connections on port `9000`
- Handles multiple clients concurrently using goroutines
- Reads incoming messages line-by-line
- Prints received messages to the server console

## 🛠 Requirements

- Go 1.18 or later (earlier versions may also work)

Check your Go version:
```bash
go version
