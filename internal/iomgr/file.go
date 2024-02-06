package iomgr

import (
	"io"
	"os"
)

type FileIOMgr struct {
	fd *os.File
}

func NewFileIOMgr(filePath string) (*FileIOMgr, error) {
	fd, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, DataFilePerm)
	if err != nil {
		return nil, err
	}
	return &FileIOMgr{fd: fd}, nil
}

func (f *FileIOMgr) Read(bs []byte, offset int64) (int, error) {
	var al int
	for al < len(bs) {
		n, err := f.fd.ReadAt(bs[al:], offset+int64(al))
		if err == io.EOF {
			return al, nil
		} else if err != nil {
			return 0, err
		}
		al = al + n
	}
	return al, nil
}

func (f *FileIOMgr) Write(bs []byte) (int, error) {
	return f.fd.Write(bs)
}

func (f *FileIOMgr) Sync() error {
	return f.fd.Sync()
}

func (f *FileIOMgr) Close() error {
	return f.fd.Close()
}

func (f *FileIOMgr) Size() (int64, error) {
	stat, err := f.fd.Stat()
	if err != nil {
		return 0, err
	}
	return stat.Size(), nil
}
