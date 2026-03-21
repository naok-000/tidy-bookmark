package bookmark

import (
	"strconv"
	"strings"
)

type Bookmark struct {
	URL string
}

func NewBookmark(url string) Bookmark {
	return Bookmark{
		URL: url,
	}
}

type BookmarkList struct {
	Items []Bookmark
}

func (l *BookmarkList) Add(url string) {
	l.Items = append(l.Items, NewBookmark(url))
}

func (l *BookmarkList) Remove(index int) {
	l.Items = append(l.Items[:index], l.Items[index+1:]...)
}

func (l *BookmarkList) Show() string {
	var result strings.Builder
	for i, item := range l.Items {
		result.WriteString(strconv.Itoa(i+1) + ". " + item.URL + "\n")
	}
	return result.String()
}
