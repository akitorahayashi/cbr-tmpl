package internal

import "fmt"

// Storage is an abstraction for item persistence.
type Storage interface {
	Add(id, content string) error
	List() ([]string, error)
	Delete(id string) error
	Exists(id string) bool
	Get(id string) (string, error)
}

// ItemNotFoundError is the error when an item is not found.
type ItemNotFoundError struct {
	ID string
}

func (e *ItemNotFoundError) Error() string {
	return fmt.Sprintf("item '%s' not found", e.ID)
}

// ItemExistsError is the error when an item already exists.
type ItemExistsError struct {
	ID string
}

func (e *ItemExistsError) Error() string {
	return fmt.Sprintf("item '%s' already exists", e.ID)
}
