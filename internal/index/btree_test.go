package index

import (
	"reflect"
	"testing"
)

type testStruct struct {
	FileID uint32
	Offset uint64
}

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
			bt := NewBTree[*testStruct]()
			bt.Put(nil, &testStruct{FileID: 1001, Offset: 1})
			bt.Put([]byte("key01"), &testStruct{FileID: 1001, Offset: 2})
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
		want *testStruct
	}{
		{"test", args{key: nil}, &testStruct{FileID: 1001, Offset: 1}},
		{"test", args{key: []byte("key01")}, &testStruct{FileID: 1001, Offset: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBTree[*testStruct]()
			bt.Put(nil, &testStruct{FileID: 1001, Offset: 1})
			bt.Put([]byte("key01"), &testStruct{FileID: 1001, Offset: 2})
			bt.Put([]byte("key01"), &testStruct{FileID: 1001, Offset: 3})
			if got, _ := bt.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBTree_Put(t *testing.T) {
	type args struct {
		key []byte
		pos *testStruct
	}
	tests := []struct {
		name string
		args args
		want *testStruct
	}{
		{"test", args{nil, &testStruct{FileID: 1001, Offset: 1}}, nil},
		{"test", args{[]byte("key01"), &testStruct{FileID: 1001, Offset: 2}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bt := NewBTree[*testStruct]()
			if got, _ := bt.Put(tt.args.key, tt.args.pos); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Put() = %v, want %v", got, tt.want)
			}
		})
	}
}
