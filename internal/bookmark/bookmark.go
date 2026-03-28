package bookmark

import (
	"errors"
	"strconv"
	"strings"
)

type Bookmark struct {
	URL string
}

type Store interface {
	Load() ([]Bookmark, error)
	Save([]Bookmark) error
}

func NewBookmark(url string) Bookmark {
	return Bookmark{
		URL: url,
	}
}

func Add(store Store, url string) error {
	bookmarks, err := store.Load()
	if err != nil {
		return err
	}

	bookmarks = append(bookmarks, NewBookmark(url))

	return store.Save(bookmarks)
}

func List(store Store) (string, error) {
	bookmarks, err := store.Load()
	if err != nil {
		return "", err
	}

	var result strings.Builder
	for i, bookmark := range bookmarks {
		result.WriteString(strconv.Itoa(i) + ". " + bookmark.URL + "\n")
	}

	return result.String(), nil
}

func Remove(store Store, index int) error {
	bookmarks, err := store.Load()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(bookmarks) {
		return errors.New("bookmark ID out of range")
	}

	bookmarks = append(bookmarks[:index], bookmarks[index+1:]...)

	return store.Save(bookmarks)
}
