package bookmark

import (
	"testing"
)

type fakeStore struct {
	loadList  BookmarkList
	loadErr   error
	saveErr   error
	savedList BookmarkList
}

func (s *fakeStore) Load() (BookmarkList, error) {
	return s.loadList, s.loadErr
}

func (s *fakeStore) Save(list BookmarkList) error {
	s.savedList = list
	return s.saveErr
}

func TestBookmarkListAdd(t *testing.T) {
	var list BookmarkList
	url0 := "https://x.com/nao_k000/status/1948052943485210629"
	url1 := "https://github.com/naok-000/tidy-bookmark/blob/main/LICENSE"

	list.Add(url0)
	list.Add(url1)

	if len(list.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(list.Items))
	}

	if list.Items[0].URL != url0 {
		t.Fatalf("expected first URL %q, got %q", url0, list.Items[0].URL)
	}

	if list.Items[1].URL != url1 {
		t.Fatalf("expected second URL %q, got %q", url1, list.Items[1].URL)
	}

}

func TestBookmarkListShow(t *testing.T) {
	var list BookmarkList
	url0 := "https://x.com/nao_k000/status/1948052943485210629"
	url1 := "https://github.com/naok-000/tidy-bookmark/blob/main/LICENSE"
	list.Add(url0)
	list.Add(url1)

	if list.Show() != "0. "+url0+"\n1. "+url1+"\n" {
		t.Fatalf("unexpected Show output:\n%s", list.Show())
	}
}

func TestBookmarkListRemove(t *testing.T) {
	var list BookmarkList
	url0 := "https://x.com/nao_k000/status/1948052943485210629"
	url1 := "https://github.com/naok-000/tidy-bookmark/blob/main/LICENSE"
	url2 := "https://go.dev"
	list.Add(url0)
	list.Add(url1)
	list.Add(url2)

	list.Remove(1)

	if len(list.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(list.Items))
	}

	if list.Items[0].URL != url0 {
		t.Fatalf("expected first URL %q, got %q", url0, list.Items[0].URL)
	}

	if list.Items[1].URL != url2 {
		t.Fatalf("expected second URL %q, got %q", url2, list.Items[1].URL)
	}
}

func TestAddLoadsAndSavesUpdatedList(t *testing.T) {
	store := &fakeStore{
		loadList: BookmarkList{
			Items: []Bookmark{
				{URL: "https://go.dev"},
			},
		},
	}

	err := Add(store, "https://example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(store.savedList.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(store.savedList.Items))
	}

	if store.savedList.Items[0].URL != "https://go.dev" {
		t.Fatalf("expected first URL %q, got %q", "https://go.dev", store.savedList.Items[0].URL)
	}

	if store.savedList.Items[1].URL != "https://example.com" {
		t.Fatalf("expected second URL %q, got %q", "https://example.com", store.savedList.Items[1].URL)
	}
}

func TestListShowsLoadedBookmarks(t *testing.T) {
	store := &fakeStore{
		loadList: BookmarkList{
			Items: []Bookmark{
				{URL: "https://go.dev"},
				{URL: "https://example.com"},
			},
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
		loadList: BookmarkList{
			Items: []Bookmark{
				{URL: "https://go.dev"},
				{URL: "https://example.com"},
				{URL: "https://golang.org"},
			},
		},
	}

	err := Remove(store, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(store.savedList.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(store.savedList.Items))
	}

	if store.savedList.Items[0].URL != "https://go.dev" {
		t.Fatalf("expected first URL %q, got %q", "https://go.dev", store.savedList.Items[0].URL)
	}

	if store.savedList.Items[1].URL != "https://golang.org" {
		t.Fatalf("expected second URL %q, got %q", "https://golang.org", store.savedList.Items[1].URL)
	}
}
