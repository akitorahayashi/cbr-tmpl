package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/akitorahayashi/cbr-tmpl/internal"
)

type mockStorage struct {
	items map[string]string
}

func newMockStorage() *mockStorage {
	return &mockStorage{items: make(map[string]string)}
}

func (m *mockStorage) Add(id, content string) error {
	if _, ok := m.items[id]; ok {
		return &internal.ItemExistsError{ID: id}
	}
	m.items[id] = content
	return nil
}

func (m *mockStorage) List() ([]string, error) {
	var ids []string
	for id := range m.items {
		ids = append(ids, id)
	}
	return ids, nil
}

func (m *mockStorage) Delete(id string) error {
	if _, ok := m.items[id]; !ok {
		return &internal.ItemNotFoundError{ID: id}
	}
	delete(m.items, id)
	return nil
}

func (m *mockStorage) Exists(id string) bool {
	_, ok := m.items[id]
	return ok
}

func (m *mockStorage) Get(id string) (string, error) {
	return m.items[id], nil
}

func TestAddCommand(t *testing.T) {
	storage := newMockStorage()
	cmd := NewRootCmd(storage)

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"add", "note1", "-c", "test content"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if storage.items["note1"] != "test content" {
		t.Error("item not added")
	}
	if !strings.Contains(buf.String(), "Added") {
		t.Error("expected success message")
	}
}

func TestAddCommand_Alias(t *testing.T) {
	storage := newMockStorage()
	cmd := NewRootCmd(storage)

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"a", "note2", "-c", "alias test"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if storage.items["note2"] != "alias test" {
		t.Error("item not added via alias")
	}
}

func TestAddCommand_DuplicateError(t *testing.T) {
	storage := newMockStorage()
	storage.items["existing"] = "old content"

	cmd := NewRootCmd(storage)
	errBuf := new(bytes.Buffer)
	cmd.SetErr(errBuf)
	cmd.SetArgs([]string{"add", "existing", "-c", "new content"})

	if err := cmd.Execute(); err == nil {
		t.Error("expected error for duplicate item")
	}

	if !strings.Contains(errBuf.String(), "already exists") {
		t.Error("expected 'already exists' error message")
	}
}

func TestListCommand(t *testing.T) {
	storage := newMockStorage()
	storage.items["item1"] = "content"

	cmd := NewRootCmd(storage)
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"list"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if !strings.Contains(buf.String(), "item1") {
		t.Error("expected item1 in output")
	}
}

func TestListCommand_Alias(t *testing.T) {
	storage := newMockStorage()
	storage.items["item1"] = "content"

	cmd := NewRootCmd(storage)
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"ls"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if !strings.Contains(buf.String(), "item1") {
		t.Error("expected item1 in output via alias")
	}
}

func TestListCommand_Empty(t *testing.T) {
	storage := newMockStorage()

	cmd := NewRootCmd(storage)
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"list"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if !strings.Contains(buf.String(), "No items found") {
		t.Error("expected 'No items found' message")
	}
}

func TestDeleteCommand(t *testing.T) {
	storage := newMockStorage()
	storage.items["to-delete"] = "content"

	cmd := NewRootCmd(storage)
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"delete", "to-delete"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if _, ok := storage.items["to-delete"]; ok {
		t.Error("item should be deleted")
	}
	if !strings.Contains(buf.String(), "Deleted") {
		t.Error("expected success message")
	}
}

func TestDeleteCommand_Alias(t *testing.T) {
	storage := newMockStorage()
	storage.items["item"] = "content"

	cmd := NewRootCmd(storage)
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"rm", "item"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if _, ok := storage.items["item"]; ok {
		t.Error("item should be deleted via alias")
	}
}

func TestDeleteCommand_NotFoundError(t *testing.T) {
	storage := newMockStorage()

	cmd := NewRootCmd(storage)
	errBuf := new(bytes.Buffer)
	cmd.SetErr(errBuf)
	cmd.SetArgs([]string{"delete", "nonexistent"})

	if err := cmd.Execute(); err == nil {
		t.Error("expected error for nonexistent item")
	}

	if !strings.Contains(errBuf.String(), "not found") {
		t.Error("expected 'not found' error message")
	}
}

func TestVersionFlag(t *testing.T) {
	storage := newMockStorage()
	cmd := NewRootCmd(storage)

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--version"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if !strings.Contains(buf.String(), version) {
		t.Errorf("expected version %s in output", version)
	}
}
