package wsl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
IsWSL checks if the current environment is Windows Subsystem for Linux (WSL) by reading specific system files that contain information about the kernel and operating system. It looks for keywords like "microsoft" or "wsl" in the contents of these files to determine if it's running in a WSL environment.

Args:

	None

Returns:

	bool: True if running in WSL, false otherwise.
*/
func IsWSL() bool {
	check := func(path string) bool {
		data, err := os.ReadFile(path)
		if err != nil {
			return false
		}
		s := strings.ToLower(string(data))
		return strings.Contains(s, "microsoft") || strings.Contains(s, "wsl")
	}
	return check("/proc/version") || check("/proc/sys/kernel/osrelease")
}

/*
FindClipPath searches for the Windows clip.exe utility in a WSL environment by looking through the /mnt directory for mounted Windows drives. It constructs potential paths to clip.exe based on common Windows drive letters and checks if the file exists. If found, it returns the path to clip.exe; otherwise, it returns an error indicating that clip.exe was not found.

Args:

	None

Returns:

	string: The path to clip.exe if found.
	error: An error if clip.exe is not found in any of the expected locations.
*/
func FindClipPath() (string, error) {
	mntDir := "/mnt"
	entries, err := os.ReadDir(mntDir)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", mntDir, err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		candidate := filepath.Join(mntDir, entry.Name(), "Windows", "System32", "clip.exe")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
	}

	return "", fmt.Errorf("clip.exe not found under %s", mntDir)
}
