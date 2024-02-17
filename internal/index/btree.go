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

func (bt *BTree[T]) Put(key []byte, val T) (oldVal T, oldExist bool) {
	it := &Item[T]{Key: key, Value: val}
	bt.lock.Lock()
	defer bt.lock.Unlock()

	old := bt.tree.ReplaceOrInsert(it)
	if old == nil {
		return
	}
	if old, ok := old.(*Item[T]); ok {
		return old.Value, true
	}

	return
}

func (bt *BTree[T]) Get(key []byte) (val T, has bool) {
	itKey := &Item[T]{Key: key}
	bt.lock.RLock()
	defer bt.lock.RUnlock()

	it := bt.tree.Get(itKey)
	if it == nil {
		return val, false
	}
	if it, ok := it.(*Item[T]); ok {
		return it.Value, true
	}

	return val, false
}

func (bt *BTree[T]) Delete(key []byte) (oldVal T, ok bool) {
	it := &Item[T]{Key: key}
	bt.lock.Lock()
	defer bt.lock.Unlock()

	old := bt.tree.Delete(it)
	if old == nil {
		return
	}
	if old, ok := old.(*Item[T]); ok {
		return old.Value, true
	}

	return
}

func (bt *BTree[T]) All() []T {
	bt.lock.RLock()
	defer bt.lock.RUnlock()

	items := make([]T, 0, bt.tree.Len())

	iterFn := func(it btree.Item) bool {
		items = append(items, it.(*Item[T]).Value)
		return true
	}
	bt.tree.Ascend(iterFn)

	return items
}

func (bt *BTree[T]) Len() int {
	bt.lock.RLock()
	defer bt.lock.RUnlock()

	return bt.tree.Len()
}
