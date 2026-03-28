package cli

import (
	"bytes"
	"path/filepath"
	"testing"

	storepkg "tidy-bookmark/internal/store"

	"github.com/spf13/cobra"
)

func executeCommand(root *cobra.Command, args ...string) (string, error) {
	var buf bytes.Buffer
	root.SetOut(&buf)
	root.SetErr(&buf)
	root.SetArgs(args)

	err := root.Execute()
	return buf.String(), err
}

func TestAddCommand(t *testing.T) {
	cmd := newRootCmd(storepkg.FileStore{Path: filepath.Join(t.TempDir(), "bookmarks.txt")})

	out, err := executeCommand(cmd, "add", "https://go.dev")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != "" {
		t.Fatalf("unexpected output: %q", out)
	}
}

func TestListCommand(t *testing.T) {
	cmd := newRootCmd(storepkg.FileStore{Path: filepath.Join(t.TempDir(), "bookmarks.txt")})

	_, err := executeCommand(cmd, "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRemoveCommand(t *testing.T) {
	cmd := newRootCmd(storepkg.FileStore{Path: filepath.Join(t.TempDir(), "bookmarks.txt")})

	_, _ = executeCommand(cmd, "add", "https://go.dev")

	out, err := executeCommand(cmd, "remove", "0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != "" {
		t.Fatalf("unexpected output: %q", out)
	}
}
