package internal

import (
	"os"
	"path/filepath"
	"sort"
)

// FilesystemStorage is a file system-based Storage implementation.
type FilesystemStorage struct {
	baseDir string
}

// NewFilesystemStorage creates a new FilesystemStorage.
// If baseDir is empty, ~/.config/cbr-tmpl/items is used.
func NewFilesystemStorage(baseDir string) *FilesystemStorage {
	if baseDir == "" {
		home, _ := os.UserHomeDir()
		baseDir = filepath.Join(home, ".config", "cbr-tmpl", "items")
	}
	os.MkdirAll(baseDir, 0755)
	return &FilesystemStorage{baseDir: baseDir}
}

func (s *FilesystemStorage) itemPath(id string) string {
	return filepath.Join(s.baseDir, id+".txt")
}

func (s *FilesystemStorage) Add(id, content string) error {
	path := s.itemPath(id)
	if _, err := os.Stat(path); err == nil {
		return &ItemExistsError{ID: id}
	}
	return os.WriteFile(path, []byte(content), 0644)
}

func (s *FilesystemStorage) List() ([]string, error) {
	entries, err := os.ReadDir(s.baseDir)
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, e := range entries {
		if !e.IsDir() && filepath.Ext(e.Name()) == ".txt" {
			ids = append(ids, e.Name()[:len(e.Name())-4])
		}
	}
	sort.Strings(ids)
	return ids, nil
}

func (s *FilesystemStorage) Delete(id string) error {
	path := s.itemPath(id)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &ItemNotFoundError{ID: id}
	}
	return os.Remove(path)
}

func (s *FilesystemStorage) Exists(id string) bool {
	_, err := os.Stat(s.itemPath(id))
	return err == nil
}

func (s *FilesystemStorage) Get(id string) (string, error) {
	data, err := os.ReadFile(s.itemPath(id))
	if os.IsNotExist(err) {
		return "", nil
	}
	return string(data), err
}
