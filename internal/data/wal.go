package data

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
)

const (
	headerCrcSize      = crc32.Size
	headerTypeSize     = 1
	headerKeyMaxSize   = binary.MaxVarintLen32
	headerValueMaxSize = binary.MaxVarintLen32
	walHeaderMaxSize   = headerCrcSize + headerTypeSize + headerKeyMaxSize + headerValueMaxSize
)

var (
	ErrInvalidBytes = errors.New("invalid bytes value, wal maybe corrupted")
	ErrInvalidCRC   = errors.New("invalid crc value, wal maybe corrupted")
)

type WALHeader struct {
	CRC        uint32
	IsDel      bool
	KeySize    int
	ValueSize  int
	HeaderSize int
}

type KVPair struct {
	Key, Value []byte
}

type WriteAheadLog struct {
	WALHeader
	KVPair
}

func (l *WriteAheadLog) Encode() ([]byte, int) {
	header := make([]byte, walHeaderMaxSize)
	writeIdx := headerCrcSize + headerTypeSize - 1
	var isDelBit byte
	if l.IsDel {
		isDelBit = 1
	}
	header[writeIdx] = isDelBit
	writeIdx += headerTypeSize
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

func (l *WriteAheadLog) DecodeHeader(bs []byte) error {
	if len(bs) < headerCrcSize+headerTypeSize {
		return ErrInvalidBytes
	}
	l.CRC = binary.LittleEndian.Uint32(bs[:headerCrcSize])
	readIdx := headerCrcSize + headerTypeSize - 1

	isDelBit := bs[readIdx]
	if isDelBit == 1 {
		l.IsDel = true
	}
	readIdx += headerTypeSize

	keySize, n := binary.Varint(bs[readIdx:])
	if n <= 0 {
		return ErrInvalidBytes
	}
	l.KeySize = int(keySize)
	readIdx += n

	valueSize, n := binary.Varint(bs[readIdx:])
	if n <= 0 {
		return ErrInvalidBytes
	}
	l.ValueSize = int(valueSize)
	readIdx += n

	l.HeaderSize = readIdx
	return nil
}

func (l *WriteAheadLog) DecodeBody(bs []byte) error {
	if len(bs) < l.KeySize+l.ValueSize {
		return ErrInvalidBytes
	}
	if l.KeySize > 0 {
		l.Key = bs[:l.KeySize]
	}
	if l.ValueSize > 0 {
		l.Value = bs[l.KeySize : l.KeySize+l.ValueSize]
	}
	return nil
}

func (l *WriteAheadLog) GetSize() int {
	return l.HeaderSize + l.KeySize + l.ValueSize
}

func (l *WriteAheadLog) CheckCRC(headerBytes []byte) bool {
	if l.HeaderSize < headerCrcSize {
		return false
	}
	if len(headerBytes) < l.HeaderSize {
		return false
	}
	crc := crc32.ChecksumIEEE(headerBytes[headerCrcSize:l.HeaderSize])
	crc = crc32.Update(crc, crc32.IEEETable, l.Key)
	crc = crc32.Update(crc, crc32.IEEETable, l.Value)

	return crc == l.CRC
}
