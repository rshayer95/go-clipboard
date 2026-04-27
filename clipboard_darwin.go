//go:build darwin

package clipboard

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

/*
darwinClipboard implements the Clipboard interface for macOS using the pbcopy command. It provides a method to copy text to the clipboard on macOS systems. The Copy method uses a mutex to ensure thread safety when accessing the clipboard.
*/
type darwinClipboard struct {
	mu sync.Mutex
}

/*
New creates a new instance of darwinClipboard, which implements the Clipboard interface for macOS. It does not require any initialization since pbcopy is expected to be available on all macOS systems.

Args:

	None

Returns:

	Clipboard: An instance of darwinClipboard that can be used to copy text to the clipboard on macOS.
	Error: Always returns nil since no initialization is needed.
*/
func New() (Clipboard, error) { return &darwinClipboard{}, nil }

/*
Copy copies the given text to the clipboard on macOS using the pbcopy command. It locks a mutex to ensure that clipboard access is thread-safe, then executes the pbcopy command with the provided text as input.

Args:

	text (string): The text to be copied to the clipboard.

Returns:

	error: An error if the pbcopy command fails to execute, otherwise nil.
*/
func (c *darwinClipboard) Copy(text string) error {
	if text == "" {
		return fmt.Errorf("cannot copy empty text to clipboard")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(text)
	return cmd.Run()
}

func (c *darwinClipboard) CopyToHost(text string) error {
	if text == "" {
		return fmt.Errorf("cannot copy empty text to clipboard")
	}
	return fmt.Errorf("CopyToHost is not supported on macOS")
}
