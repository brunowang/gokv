package index

import (
	"github.com/google/btree"
	"sync"
)

type BTree[T any] struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

func NewBTree[T any]() *BTree[T] {
	return &BTree[T]{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

func (bt *BTree[T]) Put(key []byte, val T) bool {
	it := &Item[T]{Key: key, Value: val}
	bt.lock.Lock()
	defer bt.lock.Unlock()

	bt.tree.ReplaceOrInsert(it)

	return true
}

func (bt *BTree[T]) Get(key []byte) (val T) {
	itKey := &Item[T]{Key: key}
	bt.lock.RLock()
	defer bt.lock.RUnlock()

	it := bt.tree.Get(itKey)
	if it == nil {
		return val
	}
	if it, ok := it.(*Item[T]); ok {
		return it.Value
	}

	return val
}

func (bt *BTree[T]) Delete(key []byte) bool {
	it := &Item[T]{Key: key}
	bt.lock.Lock()
	defer bt.lock.Unlock()

	return bt.tree.Delete(it) != nil
}
