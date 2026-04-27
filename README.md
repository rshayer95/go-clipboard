# go-clipboard

A cross-platform Go library for clipboard operations.  
Supports Linux, macOS, Windows, and WSL environments.

## Features

- Copy text to the native clipboard on all major platforms
- Special support for copying from WSL to Windows clipboard
- Simple, thread-safe API

## Design

Platform-specific implementations use Go build tags, so only the required clipboard backend is compiled into the final binary.

## Usage

```go
import "github.com/rshayer95/go-clipboard/clipboard"

cb, err := clipboard.New()
if err != nil {
    log.Fatal(err)
}
cb.Copy("Hello, world!")
cb.CopyToHost("Hello from WSL!")
```
