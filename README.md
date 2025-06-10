# go-for-rays
Taking on the Ray Tracer Challenge by Jamis Buck

## Dev Setup

To develop this project, you'll need:

1. **Go 1.24+** - [Install Go](https://golang.org/doc/install)
2. **golangci-lint** (for code linting) - Install with:
   ```bash
   # On Linux/macOS using binary install
   curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0

   # Or using Go install
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

   # Or using package managers:
   # brew install golangci-lint                 # macOS
   # sudo apt install golangci-lint             # Debian
   # sudo snap install golangci-lint --classic  # Ubuntu
   ```

Once you have the dependencies installed, run `make help` to see all available development targets.
