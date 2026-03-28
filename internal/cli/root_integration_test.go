package cli

import (
	"path/filepath"
	"testing"

	storepkg "tidy-bookmark/internal/store"
)

func TestAddThenListAcrossInvocations(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bookmarks.txt")

	_, err := executeCommand(newRootCmd(storepkg.FileStore{Path: path}), "add", "https://go.dev")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := executeCommand(newRootCmd(storepkg.FileStore{Path: path}), "list")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "0. https://go.dev\n"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
