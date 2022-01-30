package bytebuilder

import (
	"fmt"
	"io"
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

// ReadBytes removes the first n bytes from in and returns it.
// If the read failed, returns nil.
func ReadBytes(in *[]byte, n int) []byte {

	if len(*in) < n || n < 0 {
		return nil
	}

	out := (*in)[:n]
	*in = (*in)[n:]

	return out
}

// Skip removes the first n bytes from in.
// Returns whether it was successful.
func Skip(in *[]byte, n int) bool {
	return ReadBytes(in, n) != nil
}

// ReadUint8 removes the first byte from in and returns it as an uint8.
// The bool indicates whether the read was successful.
func ReadUint8(in *[]byte) (uint8, bool) {

	v := ReadBytes(in, 1)
	if v == nil {
		return 0, false
	}

	return uint8(v[0]), true
}

// ReadUint16 removes the first bytes from in and returns it as an uint16.
// The bool indicates whether the read was successful.
func ReadUint16(in *[]byte) (uint16, bool) {

	v := ReadBytes(in, 2)
	if v == nil {
		return 0, false
	}

	return uint16(v[0])<<8 | uint16(v[1]), true
}

// ReadUint24 removes the first bytes from in and returns it as a uint32.
// The bool indicates whether the read was successful.
func ReadUint24(in *[]byte) (uint32, bool) {

	v := ReadBytes(in, 3)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2]), true
}

// ReadUint32 removes the first bytes from in and returns it as an uint32.
// The bool indicates whether the read was successful.
func ReadUint32(in *[]byte) (uint32, bool) {

	v := ReadBytes(in, 4)
	if v == nil {
		return 0, false
	}

	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3]), true
}

// ReadReaderBytes reads n byte from in.
// At the end of the file, io.EOF is returned.
func ReadReaderBytes(in io.Reader, n int) ([]byte, error) {

	if n < 0 {
		return []byte{}, fmt.Errorf("number is less than zero")
	}

	b := make([]byte, n)

	_, err := in.Read(b)

	return b, err
}

// ReadReaderSkip skips n bytes in in.
// At the end of the file, io.EOF is returned.
func SkipReader(in io.Reader, n int) error {

	_, err := ReadReaderBytes(in, n)

	return err
}

// ReadReaderUint8 reads a byte from in and returns is as an uint8.
// At the end of the file, io.EOF is returned.
func ReadReaderUint8(in io.Reader) (uint8, error) {

	v, err := ReadReaderBytes(in, 1)

	return uint8(v[0]), err
}

// ReadReaderUint16 reads bytes from in and returns it as an uint16.
// At the end of the file io.EOF returned.
func ReadReaderUint16(in io.Reader) (uint16, error) {

	v, err := ReadReaderBytes(in, 2)
	if err != nil {
		return 0, err
	}

	return uint16(v[0])<<8 | uint16(v[1]), err
}

// ReadReaderUint24 reads bytes from in and returns it as an uint32.
// At the end of the file, io.EOF is returned.
func ReadReaderUint24(in io.Reader) (uint32, error) {

	v, err := ReadReaderBytes(in, 3)
	if err != nil {
		return 0, err
	}

	return uint32(v[0])<<16 | uint32(v[1])<<8 | uint32(v[2]), err
}

// ReadReaderUint32 reads bytes from in and returns it as an uint32.
// At the end of the file io.EOF returned.
func ReadReaderUint32(in io.Reader) (uint32, error) {

	v, err := ReadReaderBytes(in, 4)
	if err != nil {
		return 0, err
	}

	return uint32(v[0])<<24 | uint32(v[1])<<16 | uint32(v[2])<<8 | uint32(v[3]), err
}
