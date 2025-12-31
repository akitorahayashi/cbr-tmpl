package internal

import (
	"errors"
	"testing"
)

func TestFilesystemStorage_Add(t *testing.T) {
	s := NewFilesystemStorage(t.TempDir())

	if err := s.Add("item1", "content1"); err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	// 重複追加でエラー
	err := s.Add("item1", "new content")
	var existsErr *ItemExistsError
	if !errors.As(err, &existsErr) {
		t.Errorf("expected ItemExistsError, got %v", err)
	}
}

func TestFilesystemStorage_List(t *testing.T) {
	s := NewFilesystemStorage(t.TempDir())
	_ = s.Add("b", "")
	_ = s.Add("a", "")

	items, _ := s.List()
	if len(items) != 2 || items[0] != "a" || items[1] != "b" {
		t.Errorf("expected [a b], got %v", items)
	}
}

func TestFilesystemStorage_Delete(t *testing.T) {
	s := NewFilesystemStorage(t.TempDir())
	_ = s.Add("item", "content")

	if err := s.Delete("item"); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	// 存在しないアイテム削除でエラー
	err := s.Delete("nonexistent")
	var notFoundErr *ItemNotFoundError
	if !errors.As(err, &notFoundErr) {
		t.Errorf("expected ItemNotFoundError, got %v", err)
	}
}

func TestFilesystemStorage_Exists(t *testing.T) {
	s := NewFilesystemStorage(t.TempDir())

	if s.Exists("nonexistent") {
		t.Error("Exists should return false for nonexistent item")
	}

	_ = s.Add("item", "content")
	if !s.Exists("item") {
		t.Error("Exists should return true for existing item")
	}
}

func TestFilesystemStorage_Get(t *testing.T) {
	s := NewFilesystemStorage(t.TempDir())

	// 存在しないアイテム
	content, err := s.Get("nonexistent")
	if err != nil || content != "" {
		t.Errorf("Get nonexistent: expected ('', nil), got ('%s', %v)", content, err)
	}

	// 存在するアイテム
	_ = s.Add("item", "test content")
	content, err = s.Get("item")
	if err != nil || content != "test content" {
		t.Errorf("Get item: expected ('test content', nil), got ('%s', %v)", content, err)
	}
}
