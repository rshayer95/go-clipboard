package clipboard_test

import (
	"errors"
	"os"
	"runtime"
	"testing"

	"github.com/rshayer95/go-clipboard"
)

func TestClipboardInterface(t *testing.T) {
	cb, err := clipboard.New()
	if err != nil {
		t.Fatalf("failed to create clipboard: %v", err)
	}

	// Test Copy with empty string
	err = cb.Copy("")
	if err == nil {
		t.Error("expected error when copying empty string, got nil")
	}

	// Test CopyToHost with empty string
	err = cb.CopyToHost("")
	if err == nil {
		t.Error("expected error when copying empty string to host, got nil")
	}
}

func TestClipboardCopy(t *testing.T) {
	cb, err := clipboard.New()
	if err != nil {
		t.Skipf("clipboard not available: %v", err)
	}

	text := "Hello, Clipboard!"
	err = cb.Copy(text)
	if err != nil {
		t.Errorf("Copy failed: %v", err)
	}
}

func TestClipboardCopyToHost(t *testing.T) {
	cb, err := clipboard.New()
	if err != nil {
		t.Skipf("clipboard not available: %v", err)
	}

	if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
		err = cb.CopyToHost("test")
		if err == nil {
			t.Error("expected error for CopyToHost on unsupported OS, got nil")
		}
	} else {
		// On Linux, CopyToHost may succeed if WSL/clip.exe is available, otherwise error
		err = cb.CopyToHost("test")
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			t.Logf("CopyToHost failed as expected: %v", err)
		}
	}
}
