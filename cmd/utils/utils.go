package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gliderlabs/ssh"
)

func LoadFilesFromDir(dir string) ([]string, error) {
	var files []string

	// Walk through the directory to gather file names.
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a regular file (not a directory).
		if !info.IsDir() {
			// Get the file name without extension.
			filename := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			files = append(files, filename)
		}
		return nil
	})

	return files, err
}

func ReadFileContent(path string) (string, error) {
	data, err := os.ReadFile(path) // Read the file content
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GetTerminalColorSupport(s ssh.Session) int {
	// Get the TERM environment variable from the SSH session
	termEnv := os.Getenv("TERM")
	fmt.Printf("Client TERM: %s\n", termEnv)

	// Get terminal info using TERM and terminfo library
	// This will give us detailed information about the terminal capabilities
	switch termEnv {
	case "xterm-256color", "screen-256color":
		return 256
	case "xterm", "screen":
		return 16
	case "xterm-16color", "screen-16color":
		return 16
	case "xterm-88color", "screen-88color":
		return 88
	default:
		return 8 // Assume basic 8-color support as the fallback
	}
}
