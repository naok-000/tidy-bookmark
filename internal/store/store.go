package store

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"tidy-bookmark/internal/bookmark"
)

type FileStore struct {
	Path string
}

func (s FileStore) Load() (bookmark.BookmarkList, error) {
	file, err := os.Open(s.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return bookmark.BookmarkList{}, nil
		}
		return bookmark.BookmarkList{}, err
	}

	var list bookmark.BookmarkList
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list.Add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return bookmark.BookmarkList{}, errors.Join(err, file.Close())
	}

	if err := file.Close(); err != nil {
		return bookmark.BookmarkList{}, err
	}

	return list, nil
}

func (s FileStore) Save(list bookmark.BookmarkList) error {
	var content strings.Builder
	for _, item := range list.Items {
		content.WriteString(item.URL)
		content.WriteString("\n")
	}

	return os.WriteFile(s.Path, []byte(content.String()), 0o644)
}
