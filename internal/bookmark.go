package bookmark

type Bookmark struct {
	URL string
}

type BookmarkList struct {
	Items []Bookmark
}

func (l *BookmarkList) Add(url string) {
	l.Items = append(l.Items, Bookmark{URL: url})
}
