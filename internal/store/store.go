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

func (s FileStore) Load() ([]bookmark.Bookmark, error) {
	file, err := os.Open(s.Path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	var bookmarks []bookmark.Bookmark
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bookmarks = append(bookmarks, bookmark.NewBookmark(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Join(err, file.Close())
	}

	if err := file.Close(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}

func (s FileStore) Save(bookmarks []bookmark.Bookmark) error {
	var content strings.Builder
	for _, item := range bookmarks {
		content.WriteString(item.URL)
		content.WriteString("\n")
	}

	return os.WriteFile(s.Path, []byte(content.String()), 0o644)
}
