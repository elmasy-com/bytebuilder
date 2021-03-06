package bytebuilder

import (
	"time"
)

// ReadBytes removes the first n bytes from b and returns it.
// If the read failed, returns nil.
func (b *Buffer) ReadBytes(n int) []byte {

	if len(b.b) < n || n < 0 {
		return nil
	}

	v := b.b[:n]
	b.b = b.b[n:]

	return v
}

// Skip removes the first n bytes from b.
// Returns whether it was successful.
func (b *Buffer) Skip(n int) bool {
	return b.ReadBytes(n) != nil
}

// ReadUint8 removes the first byte from b and returns it as an uint8.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint8() (uint8, bool) {

	v := b.ReadBytes(1)
	if v == nil {
		return 0, false
	}

	return uint8(v[0]), true
}

// ReadUint16 removes the first bytes from b and returns it as an uint16.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint16() (uint16, bool) {

	v := b.ReadBytes(2)
	if v == nil {
		return 0, false
	}

	return uint16(v[0])<<8 | uint16(v[1]), true
}

// ReadUint24 removes the first bytes from b and returns it as a uint32.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint24() (uint32, bool) {

	v := b.ReadBytes(3)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2]), true
}

// ReadUint32 removes the first bytes from b and returns it as an uint32.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadUint32() (uint32, bool) {

	v := b.ReadBytes(4)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3]), true
}

// ReadInt removes the first bytes (depends on bitSize) from b and returns it as an int.
func (b *Buffer) ReadInt(bitSize int) (int, bool) {

	switch bitSize {
	case 8:
		l, ok := b.ReadUint8()
		return int(l), ok
	case 16:
		l, ok := b.ReadUint16()
		return int(l), ok
	case 24:
		l, ok := b.ReadUint24()
		return int(l), ok
	case 32:
		l, ok := b.ReadUint32()
		return int(l), ok
	default:
		panic("invalid bitSize value")
	}
}

// ReadGMTUnixTime32 removes the first bytes from b and returns it as an unix time.
// The bool indicates whether the read was successful.
func (b *Buffer) ReadGMTUnixTime32() (time.Time, bool) {

	v, ok := b.ReadUint32()
	if !ok {
		return time.Time{}, false
	}

	return time.Unix(int64(v), 0), true
}

// ReadVector reads the length of bytes then the bytes itself.
// The length type is depend on bitSize (eg.: uint8, uint16, uint24, uint32).
// Therefore, bitSize must be 8/16/24/32.
// If bitSize is an invalid number, this function panics.
func (b *Buffer) ReadVector(bitSize int) ([]byte, bool) {

	n, ok := b.ReadInt(bitSize)
	if !ok {
		return []byte{}, false
	}

	v := b.ReadBytes(n)
	if v == nil {
		return []byte{}, false
	}

	return v, true
}
