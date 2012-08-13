package spiderDB

import "fmt"

type KeyNotFoundError struct { Key string }

type dbError struct {
	ErrStr string
	}

func (e *dbError) Error() string {
	return fmt.Sprintf("%s", e.ErrStr)
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("%s Not found", e.Key)
}
