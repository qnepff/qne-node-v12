# QNE Node - QUIC Server

A cross-platform QUIC/HTTP3 server implementation for the QNE Community Network.

## Features

- HTTP/3 (QUIC) server with HTTP/2 fallback
- Automatic protocol negotiation
- Graceful shutdown handling
- Self-signed TLS certificate generation
- Simple test frontend
- Systemd service integration
- Cross-platform support via goreleaser

## Requirements

- Go 1.21 or later
- For development: goreleaser

## Building

```bash
go build
```

## Development

Run the server locally:

```bash
go run main.go
```

Visit https://localhost:4444 in your browser to test the connection.

## Production Deployment

1. Build the release:
```bash
goreleaser build --snapshot --clean
```

2. Install the systemd service:
```bash
sudo cp qne-node.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable qne-node
sudo systemctl start qne-node
```

## License

Private - All rights reserved
