package bookmark

import (
	"testing"
)

type fakeStore struct {
	loadList  []Bookmark
	loadErr   error
	saveErr   error
	savedList []Bookmark
}

func (s *fakeStore) Load() ([]Bookmark, error) {
	return s.loadList, s.loadErr
}

func (s *fakeStore) Save(list []Bookmark) error {
	s.savedList = list
	return s.saveErr
}

func TestAddLoadsAndSavesUpdatedList(t *testing.T) {
	store := &fakeStore{
		loadList: []Bookmark{
			{URL: "https://go.dev"},
		},
	}

	err := Add(store, "https://example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(store.savedList) != 2 {
		t.Fatalf("expected 2 items, got %d", len(store.savedList))
	}

	if store.savedList[0].URL != "https://go.dev" {
		t.Fatalf("expected first URL %q, got %q", "https://go.dev", store.savedList[0].URL)
	}

	if store.savedList[1].URL != "https://example.com" {
		t.Fatalf("expected second URL %q, got %q", "https://example.com", store.savedList[1].URL)
	}
}

func TestListShowsLoadedBookmarks(t *testing.T) {
	store := &fakeStore{
		loadList: []Bookmark{
			{URL: "https://go.dev"},
			{URL: "https://example.com"},
		},
	}

	got, err := List(store)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := "0. https://go.dev\n1. https://example.com\n"
	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestRemoveLoadsAndSavesUpdatedList(t *testing.T) {
	store := &fakeStore{
		loadList: []Bookmark{
			{URL: "https://go.dev"},
			{URL: "https://example.com"},
			{URL: "https://golang.org"},
		},
	}

	err := Remove(store, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(store.savedList) != 2 {
		t.Fatalf("expected 2 items, got %d", len(store.savedList))
	}

	if store.savedList[0].URL != "https://go.dev" {
		t.Fatalf("expected first URL %q, got %q", "https://go.dev", store.savedList[0].URL)
	}

	if store.savedList[1].URL != "https://golang.org" {
		t.Fatalf("expected second URL %q, got %q", "https://golang.org", store.savedList[1].URL)
	}
}
