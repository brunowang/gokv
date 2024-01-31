package index

import (
	"bytes"
	"github.com/google/btree"
)

type Indexer[T any] interface {
	Put(item Item[T]) bool
	Get(key []byte) T
	Delete(key []byte) bool
}

type Item[T any] struct {
	Key   []byte
	Value T
}

func (a *Item[T]) Less(i btree.Item) bool {
	if i, ok := i.(*Item[T]); ok {
		return bytes.Compare(a.Key, i.Key) == -1
	}
	return false
}

type FilePos struct {
	FileID uint32
	Offset uint64
}
