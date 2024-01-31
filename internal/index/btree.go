package index

import (
	"github.com/google/btree"
	"sync"
)

type BTree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

func NewBTree() *BTree {
	return &BTree{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

func (bt *BTree) Put(key []byte, pos *FilePos) bool {
	it := &Item[*FilePos]{Key: key, Value: pos}
	bt.lock.Lock()
	defer bt.lock.Unlock()

	bt.tree.ReplaceOrInsert(it)

	return true
}

func (bt *BTree) Get(key []byte) *FilePos {
	itKey := &Item[*FilePos]{Key: key}
	bt.lock.RLock()
	defer bt.lock.RUnlock()

	it := bt.tree.Get(itKey)
	if it == nil {
		return nil
	}
	if it, ok := it.(*Item[*FilePos]); ok {
		return it.Value
	}

	return nil
}

func (bt *BTree) Delete(key []byte) bool {
	it := &Item[*FilePos]{Key: key}
	bt.lock.Lock()
	defer bt.lock.Unlock()

	return bt.tree.Delete(it) != nil
}
