package bookmark

import (
	"testing"

	"github.com/google/uuid"
)

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

	if list.Items[0].id == uuid.Nil {
		t.Fatal("expected non-empty ID for first bookmark")
	}

	if list.Items[1].URL != url1 {
		t.Fatalf("expected second URL %q, got %q", url1, list.Items[1].URL)
	}

	if list.Items[1].id == uuid.Nil {
		t.Fatal("expected non-empty ID for second bookmark")
	}
}

func TestBookmarkListShow(t *testing.T) {
	var list BookmarkList
	url0 := "https://x.com/nao_k000/status/1948052943485210629"
	url1 := "https://github.com/naok-000/tidy-bookmark/blob/main/LICENSE"
	list.Add(url0)
	list.Add(url1)

	if list.Show() != "1. "+url0+"\n2. "+url1+"\n" {
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
