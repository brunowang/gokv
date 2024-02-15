package index

import (
	"github.com/brunowang/gokv/internal/data"
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
			bt := NewBTree[*data.FilePos]()
			bt.Put(nil, &data.FilePos{FileID: 1001, Offset: 1})
			bt.Put([]byte("key01"), &data.FilePos{FileID: 1001, Offset: 2})
			if _, got := bt.Delete(tt.args.key); got != tt.want {
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
		want *data.FilePos
	}{
		{"test", args{key: nil}, &data.FilePos{FileID: 1001, Offset: 1}},
		{"test", args{key: []byte("key01")}, &data.FilePos{FileID: 1001, Offset: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBTree[*data.FilePos]()
			bt.Put(nil, &data.FilePos{FileID: 1001, Offset: 1})
			bt.Put([]byte("key01"), &data.FilePos{FileID: 1001, Offset: 2})
			bt.Put([]byte("key01"), &data.FilePos{FileID: 1001, Offset: 3})
			if got := bt.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Put(t *testing.T) {
	type args struct {
		key []byte
		pos *data.FilePos
	}
	tests := []struct {
		name string
		args args
		want *data.FilePos
	}{
		{"test", args{nil, &data.FilePos{FileID: 1001, Offset: 1}}, nil},
		{"test", args{[]byte("key01"), &data.FilePos{FileID: 1001, Offset: 2}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBTree[*data.FilePos]()
			if got := bt.Put(tt.args.key, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}
