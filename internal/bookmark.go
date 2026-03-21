package bookmark

import "github.com/google/uuid"

type Bookmark struct {
	id  uuid.UUID
	URL string
}

func NewBookmark(url string) Bookmark {
	return Bookmark{
		id:  uuid.New(),
		URL: url,
	}
}

type BookmarkList struct {
	Items []Bookmark
}

func (l *BookmarkList) Add(url string) {
	l.Items = append(l.Items, NewBookmark(url))
}
