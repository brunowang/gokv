package iomgr

import (
	"os"
	"testing"
)

const testFilePath = "./test.data"

func TestFileIOMgr_Close(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := NewFileIOMgr(testFilePath)
			defer os.RemoveAll(testFilePath)
			if err := f.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileIOMgr_Read(t *testing.T) {
	type args struct {
		bs     []byte
		offset int64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"test", args{make([]byte, 4), 0}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := NewFileIOMgr(testFilePath)
			defer os.RemoveAll(testFilePath)
			f.Write([]byte(tt.name))
			got, err := f.Read(tt.args.bs, tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileIOMgr_Size(t *testing.T) {
	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		{"test", 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := NewFileIOMgr(testFilePath)
			defer os.RemoveAll(testFilePath)
			f.Write([]byte(tt.name))
			got, err := f.Size()
			if (err != nil) != tt.wantErr {
				t.Errorf("Size() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Size() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileIOMgr_Sync(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := NewFileIOMgr(testFilePath)
			defer os.RemoveAll(testFilePath)
			if err := f.Sync(); (err != nil) != tt.wantErr {
				t.Errorf("Sync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileIOMgr_Write(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"test", args{[]byte("")}, 0, false},
		{"test", args{[]byte("test")}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := NewFileIOMgr(testFilePath)
			defer os.RemoveAll(testFilePath)
			got, err := f.Write(tt.args.bs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Write() got = %v, want %v", got, tt.want)
			}
		})
	}
}
