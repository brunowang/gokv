package index

import (
	"bytes"
	"github.com/google/btree"
)

type Indexer[T any] interface {
	Put(key []byte, val T) (T, bool)
	Get(key []byte) (T, bool)
	Delete(key []byte) (T, bool)
	All() []T
	Len() int
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
