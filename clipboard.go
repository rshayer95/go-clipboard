package clipboard

/*
Clipboard defines an interface for copying text to the clipboard across different operating systems. It includes a Copy method for copying text to the clipboard and a CopyToWindowsFromWsl method for copying text to the Windows clipboard from a WSL environment. Implementations of this interface are provided for Linux, macOS, and Windows, allowing for seamless clipboard operations across these platforms.
*/
type Clipboard interface {
	Copy(text string) error
	CopyToHost(text string) error
}
