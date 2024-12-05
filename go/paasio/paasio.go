package paasio

import (
	"io"
	"sync"
)

// Define readCounter and writeCounter types here.
type readCounter struct {
	sync.RWMutex
	reader io.Reader
	ops    int
	bytes  int64
}
type writeCounter struct {
	sync.RWMutex
	writer io.Writer
	ops    int
	bytes  int64
}

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.
type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	rc, wc := NewReadCounter(readwriter), NewWriteCounter(readwriter)
	return &readWriteCounter{rc, wc}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.Lock()
	defer rc.Unlock()

	n, err := rc.reader.Read(p)
	rc.ops++
	rc.bytes += int64(n)
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.Lock()
	defer rc.Unlock()

	return rc.bytes, rc.ops
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.Lock()
	defer wc.Unlock()

	n, err := wc.writer.Write(p)
	wc.ops++
	wc.bytes += int64(n)
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.Lock()
	defer wc.Unlock()

	return wc.bytes, wc.ops
}
