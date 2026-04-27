package clipboard

/*
Clipboard defines an interface for copying text to the clipboard across different operating systems. It includes a Copy method for copying text to the clipboard and a CopyToHost method for copying text to the Host clipboard from a WSL environment. Implementations of this interface are provided for Linux, macOS, and Windows, allowing for seamless clipboard operations across these platforms.
*/
type Clipboard interface {
	Copy(text string) error
	CopyToHost(text string) error
}
