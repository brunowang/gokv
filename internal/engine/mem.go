package engine

import "sync"

var (
	MemStore sync.Map
)

func Set(key string, value string) {
	MemStore.Store(key, value)
}
func Get(key string) string {
	if v, ok := MemStore.Load(key); ok {
		return v.(string)
	}
	return ""
}
