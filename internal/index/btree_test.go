package index

import (
	"reflect"
	"testing"
)

func TestBTree_Delete(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test", args{key: nil}, true},
		{"test", args{key: []byte("key01")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBTree()
			bt.Put(nil, &FilePos{FileID: 1001, Offset: 1})
			bt.Put([]byte("key01"), &FilePos{FileID: 1001, Offset: 2})
			if got := bt.Delete(tt.args.key); got != tt.want {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Get(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name string
		args args
		want *FilePos
	}{
		{"test", args{key: nil}, &FilePos{FileID: 1001, Offset: 1}},
		{"test", args{key: []byte("key01")}, &FilePos{FileID: 1001, Offset: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBTree()
			bt.Put(nil, &FilePos{FileID: 1001, Offset: 1})
			bt.Put([]byte("key01"), &FilePos{FileID: 1001, Offset: 2})
			bt.Put([]byte("key01"), &FilePos{FileID: 1001, Offset: 3})
			if got := bt.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Put(t *testing.T) {
	type args struct {
		key []byte
		pos *FilePos
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test", args{nil, &FilePos{FileID: 1001, Offset: 1}}, true},
		{"test", args{[]byte("key01"), &FilePos{FileID: 1001, Offset: 2}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBTree()
			if got := bt.Put(tt.args.key, tt.args.pos); got != tt.want {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}
