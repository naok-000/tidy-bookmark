package bookmark

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

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

func (l *BookmarkList) Show() string {
	var result strings.Builder
	for i, item := range l.Items {
		result.WriteString(strconv.Itoa(i+1) + ". " + item.URL + "\n")
	}
	return result.String()
}
