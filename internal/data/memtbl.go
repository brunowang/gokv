package data

import (
	"github.com/brunowang/gokv/internal/index"
	"sync/atomic"
)

type MemTable struct {
	data index.Indexer[*KVPair]
	size atomic.Uint64
}

func NewMemTable() *MemTable {
	return &MemTable{
		data: index.NewBTree[*KVPair](),
	}
}

func (m *MemTable) Put(key, value []byte) {
	m.data.Put(key, &KVPair{Key: key, Value: value})
	m.size.Add(uint64(len(key) + len(value)))
}

func (m *MemTable) Get(key []byte) ([]byte, bool) {
	val, has := m.data.Get(key)
	if !has {
		return nil, false
	}
	return val.Value, true
}

func (m *MemTable) All() []*KVPair {
	return m.data.All()
}

func (m *MemTable) Size() uint64 {
	return m.size.Load()
}

func (m *MemTable) EntriesCnt() int {
	return m.data.Len()
}
