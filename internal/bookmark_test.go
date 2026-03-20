package bookmark

import "testing"

func TestBookmarkListAdd(t *testing.T) {
	var list BookmarkList
	var urls = []string{
		"https://x.com/nao_k000/status/1948052943485210629",
		"https://github.com/naok-000/tidy-bookmark/blob/main/LICENSE",
	}

	list.Add(urls[0])
	list.Add(urls[1])

	if len(list.Items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(list.Items))
	}

	if list.Items[0].URL != urls[0] {
		t.Fatalf("expected first URL %q, got %q", urls[0], list.Items[0].URL)
	}

	if list.Items[1].URL != urls[1] {
		t.Fatalf("expected second URL %q, got %q", urls[1], list.Items[1].URL)
	}
}
