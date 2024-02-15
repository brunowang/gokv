package data

import (
	"errors"
	"fmt"
	"github.com/brunowang/gokv/internal/iomgr"
	"path/filepath"
)

const (
	WalFileSuffix = ".wal"
)

var (
	ErrInvalidCRC = errors.New("invalid crc value, wal maybe corrupted")
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
