package store

import (
	"os"
	"path/filepath"
	"testing"

	"tidy-bookmark/internal/bookmark"
)

func TestLoadNotExistReturnsEmptyList(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bookmarks.txt")
	store := FileStore{Path: path}

	list, err := store.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 0 {
		t.Fatalf("expected empty list, got %d items", len(list))
	}
}

func TestSaveWritesOneURLPerLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bookmarks.txt")
	store := FileStore{Path: path}
	list := []bookmark.Bookmark{
		{URL: "https://go.dev"},
		{URL: "https://example.com"},
	}

	if err := store.Save(list); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "https://go.dev\nhttps://example.com\n"
	if string(got) != want {
		t.Fatalf("unexpected file contents:\n%s", string(got))
	}
}

func TestLoadReadsBookmarksInOrder(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bookmarks.txt")
	err := os.WriteFile(path, []byte("https://go.dev\nhttps://example.com\n"), 0o644)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	store := FileStore{Path: path}

	list, err := store.Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(list) != 2 {
		t.Fatalf("expected 2 items, got %d", len(list))
	}

	if list[0].URL != "https://go.dev" {
		t.Fatalf("expected first URL %q, got %q", "https://go.dev", list[0].URL)
	}

	if list[1].URL != "https://example.com" {
		t.Fatalf("expected second URL %q, got %q", "https://example.com", list[1].URL)
	}
}

func TestSaveOverwritesExistingContent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "bookmarks.txt")
	err := os.WriteFile(path, []byte("https://old.example.com\n"), 0o644)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	store := FileStore{Path: path}
	list := []bookmark.Bookmark{
		{URL: "https://new.example.com"},
	}

	if err := store.Save(list); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "https://new.example.com\n"
	if string(got) != want {
		t.Fatalf("unexpected file contents:\n%s", string(got))
	}
}
