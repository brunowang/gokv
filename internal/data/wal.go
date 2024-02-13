package data

import (
	"encoding/binary"
	"hash/crc32"
)

const (
	headerCrcSize    = 4
	headerTypeSize   = 1
	headerKeySize    = binary.MaxVarintLen32
	headerValueSize  = binary.MaxVarintLen32
	walHeaderMaxSize = headerCrcSize + headerTypeSize + headerKeySize + headerValueSize
)

type WriteAheadLog struct {
	Key   []byte
	Value []byte
	IsDel bool
}

func (l *WriteAheadLog) Encode() ([]byte, int) {
	header := make([]byte, walHeaderMaxSize)
	writeIdx := headerCrcSize + headerTypeSize - 1
	var isDelBit byte
	if l.IsDel {
		isDelBit = 1
	}
	header[writeIdx] = isDelBit
	writeIdx += binary.PutVarint(header[writeIdx:], int64(len(l.Key)))
	writeIdx += binary.PutVarint(header[writeIdx:], int64(len(l.Value)))

	size := writeIdx + len(l.Key) + len(l.Value)
	encBytes := make([]byte, size)

	copy(encBytes[:writeIdx], header[:writeIdx])
	copy(encBytes[writeIdx:], l.Key)
	writeIdx += len(l.Key)
	copy(encBytes[writeIdx:], l.Value)

	crc := crc32.ChecksumIEEE(encBytes[headerCrcSize:])
	binary.LittleEndian.PutUint32(encBytes[:headerCrcSize], crc)

	return encBytes, size
}
