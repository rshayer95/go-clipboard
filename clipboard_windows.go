//go:build windows

package clipboard

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

type windowsClipboard struct {
	bin string
	mu  sync.Mutex
}

func New() (Clipboard, error) {
	windir := os.Getenv("WINDIR")
	if windir == "" {
		windir = os.Getenv("SystemRoot")
	}
	if windir == "" {
		return nil, fmt.Errorf("Cannot locate Windows directory: WINDIR and SystemRoot environment variables are not set")
	}
	clip := filepath.Join(windir, "System32", "clip.exe")
	if _, err := os.Stat(clip); os.IsNotExist(err) {
		return nil, fmt.Errorf("clip.exe not found at %s: %w", clip, err)
	}
	return &windowsClipboard{bin: clip}, nil
}

func (c *windowsClipboard) Copy(text string) error {
	if text == "" {
		return fmt.Errorf("cannot copy empty text to clipboard")
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	cmd := exec.Command(c.bin)
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func (c *windowsClipboard) CopyToHost(text string) error {
	if text == "" {
		return fmt.Errorf("cannot copy empty text to clipboard")
	}
	return fmt.Errorf("CopyToHost is not supported on Windows")
}
