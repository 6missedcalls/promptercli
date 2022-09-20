//go:build windows
// +build windows

package promptercli

// source: https://msdn.microsoft.com/en-us/library/aa243025(v=vs.60).aspx

var (
	// KeyBackspace is the default key for deleting input text inside a command line prompt.
	KeyBackspace rune = 8
)
