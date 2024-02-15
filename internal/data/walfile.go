package data

import (
	"fmt"
	"github.com/brunowang/gokv/internal/iomgr"
	"path/filepath"
)

const (
	WalFileSuffix = ".wal"
)

type FilePos struct {
	FileID uint32
	Offset uint64
}

type WalFile struct {
	FilePos
	ioMgr iomgr.IOManager
}

func CalcWalFilePath(dirPath string, fileID uint32) string {
	return filepath.Join(dirPath, fmt.Sprintf("%09d", fileID)+WalFileSuffix)
}

func OpenWalFile(dirPath string, fileID uint32) (*WalFile, error) {
	filePath := CalcWalFilePath(dirPath, fileID)
	fileMgr, err := iomgr.NewFileIOMgr(filePath)
	if err != nil {
		return nil, err
	}
	return &WalFile{
		FilePos: FilePos{
			FileID: fileID,
			Offset: 0,
		},
		ioMgr: fileMgr,
	}, nil
}

func (f *WalFile) Read(offset int64) (*WriteAheadLog, error) {
	fileSize, err := f.ioMgr.Size()
	if err != nil {
		return nil, err
	}
	walData := &WriteAheadLog{}

	var headerSize int64 = walHeaderMaxSize
	if offset+headerSize > fileSize {
		headerSize = fileSize - offset
	}

	headerBytes := make([]byte, headerSize)
	if _, err := f.ioMgr.Read(headerBytes, offset); err != nil {
		return nil, err
	}
	if err := walData.DecodeHeader(headerBytes); err != nil {
		return nil, err
	}
	offset += int64(walData.HeaderSize)

	bodySize := walData.GetSize() - walData.HeaderSize

	bodyBytes := make([]byte, bodySize)
	if _, err := f.ioMgr.Read(bodyBytes, offset); err != nil {
		return nil, err
	}
	if err := walData.DecodeBody(bodyBytes); err != nil {
		return nil, err
	}
	offset += int64(bodySize)

	if !walData.CheckCRC(headerBytes) {
		return nil, ErrInvalidCRC
	}

	return walData, nil
}
