package bookmark

import (
	"strconv"
	"strings"
)

type Bookmark struct {
	URL string
}

type Store interface {
	Load() (BookmarkList, error)
	Save(BookmarkList) error
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

func Add(store Store, url string) error {
	list, err := store.Load()
	if err != nil {
		return err
	}

	list.Add(url)

	return store.Save(list)
}

func List(store Store) (string, error) {
	list, err := store.Load()
	if err != nil {
		return "", err
	}

	return list.Show(), nil
}

func (l *BookmarkList) Remove(index int) {
	l.Items = append(l.Items[:index], l.Items[index+1:]...)
}

func Remove(store Store, index int) error {
	list, err := store.Load()
	if err != nil {
		return err
	}

	list.Remove(index)

	return store.Save(list)
}

func (l *BookmarkList) Show() string {
	var result strings.Builder
	for i, item := range l.Items {
		result.WriteString(strconv.Itoa(i) + ". " + item.URL + "\n")
	}
	return result.String()
}
