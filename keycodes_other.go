//go:build !windows
// +build !windows

package promptercli

import "github.com/chzyer/readline"

var (
	// KeyBackspace is the default key for deleting input text.
	KeyBackspace rune = readline.CharBackspace
)
