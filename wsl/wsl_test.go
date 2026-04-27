package wsl_test

import (
	"os"
	"testing"

	"github.com/rshayer95/go-clipboard/wsl"
)

func TestIsWSL(t *testing.T) {
	// This test is environment-dependent; just ensure it doesn't panic
	_ = wsl.IsWSL()
}

func TestFindClipPath(t *testing.T) {
	if !wsl.IsWSL() {
		t.Skip("not running in WSL environment")
	}
	path, err := wsl.FindClipPath()
	if err != nil {
		t.Logf("clip.exe not found: %v", err)
	} else {
		if _, err := os.Stat(path); err != nil {
			t.Errorf("clip.exe path returned but file does not exist: %v", err)
		}
	}
}
