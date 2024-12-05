package circular

import "fmt"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	buffer     []byte
	readIndex  int
	writeIndex int
	size       int
}

func NewBuffer(size int) *Buffer {
	bytes := make([]byte, size)
	return &Buffer{buffer: bytes}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.size == 0 {
		return 0, fmt.Errorf("buffer empty")
	}
	res := b.buffer[b.readIndex]
	b.size--
	b.readIndex++
	if b.readIndex == len(b.buffer) {
		b.readIndex = 0
	}

	return res, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.size == len(b.buffer) {
		return fmt.Errorf("buffer full")
	} else {
		b.buffer[b.writeIndex] = c
		b.size++
		b.writeIndex++
		if b.writeIndex == len(b.buffer) {
			b.writeIndex = 0
		}
		return nil
	}
}

func (b *Buffer) Overwrite(c byte) {
	if b.size < len(b.buffer) {
		b.WriteByte(c)
	} else {
		b.buffer[b.writeIndex] = c
		b.writeIndex++
		if b.writeIndex == len(b.buffer) {
			b.writeIndex = 0
		}
		b.readIndex++
		if b.readIndex == len(b.buffer) {
			b.readIndex = 0
		}
	}
}

func (b *Buffer) Reset() {
	b.buffer = make([]byte, len(b.buffer))
	b.readIndex, b.writeIndex, b.size = 0, 0, 0
}
