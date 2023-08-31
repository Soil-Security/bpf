package bpfevents

import (
	"encoding/binary"
	"errors"

	"golang.org/x/sys/unix"
)

type Decoder struct {
	ByteOrder binary.ByteOrder
}

func (e *Decoder) Byte(buf []byte, off int) (byte, int, error) {
	if off+1 > len(buf) {
		return 0, off, errors.New("overflow unpacking byte")
	}
	return buf[off], off + 1, nil
}

func (e *Decoder) Uint32(buf []byte, off int) (uint32, int, error) {
	if off+4 > len(buf) {
		return 0, off, errors.New("overflow unpacking uint32")
	}
	u := e.ByteOrder.Uint32(buf[off : off+4])
	return u, off + 4, nil
}

func (e *Decoder) Uint32AsInt(buf []byte, off int) (int, int, error) {
	i, o, err := e.Uint32(buf, off)
	return int(i), o, err
}

func (e *Decoder) Str(buf []byte, off, sz int) (string, int, error) {
	if off+sz > len(buf) {
		return "", off, errors.New("overflow unpacking string")
	}
	s := make([]byte, sz)
	_ = copy(s, buf[off:off+sz])
	return unix.ByteSliceToString(s), off + sz, nil
}
