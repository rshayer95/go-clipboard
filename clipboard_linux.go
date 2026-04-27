//go:build linux

package clipboard

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"

	"github.com/rshayer95/go-clipboard/wsl"
)

type linuxClipboard struct {
	bin    string
	args   []string
	wslBin string
	mu     sync.Mutex
}

func New() (Clipboard, error) {
	cb := &linuxClipboard{}

	// Find Linux clipboard utility
	for _, b := range []struct {
		bin  string
		args []string
	}{
		{"wl-copy", nil},
		{"xclip", []string{"-selection", "clipboard"}},
		{"xsel", []string{"--clipboard", "--input"}},
	} {
		if path, err := exec.LookPath(b.bin); err == nil {
			cb.bin = path
			cb.args = b.args
			break
		}
	}

	// Find Windows clip.exe if in WSL
	if wsl.IsWSL() {
		if clip, err := wsl.FindClipPath(); err == nil {
			cb.wslBin = clip
		}
	}

	if cb.bin == "" && cb.wslBin == "" {
		return nil, fmt.Errorf("no clipboard utility found")
	}

	return cb, nil
}

func (c *linuxClipboard) Copy(text string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.bin == "" {
		return fmt.Errorf("no Linux clipboard utility found (install wl-clipboard, xclip, or xsel)")
	}
	cmd := exec.Command(c.bin, c.args...)
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func (c *linuxClipboard) CopyToHost(text string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.wslBin == "" {
		return fmt.Errorf("CopyToHost requires WSL environment with clip.exe available")
	}
	cmd := exec.Command(c.wslBin)
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}
