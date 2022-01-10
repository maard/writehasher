// This type is an io.MultiWriter with hash.Hash[,32,64].
// After writes to an instance of WriteHasher are complete,
// .Sum(), .Sum32() or .Sum64() can be used to retrieve hashes of the writted data.

package writehasher

import (
	"hash"
	"io"
)

type WriteHasher struct {
	w   io.Writer
	h   hash.Hash
	h32 hash.Hash32
	h64 hash.Hash64
}

func NewWriteHasher(w io.Writer, h hash.Hash) *WriteHasher {
	return &WriteHasher{
		h: h,
		w: io.MultiWriter(w, h),
	}
}

func NewWriteHasher32(w io.Writer, h32 hash.Hash32) *WriteHasher {
	return &WriteHasher{
		h32: h32,
		w:   io.MultiWriter(w, h32),
	}
}

func NewWriteHasher64(w io.Writer, h64 hash.Hash64) *WriteHasher {
	return &WriteHasher{
		h64: h64,
		w:   io.MultiWriter(w, h64),
	}
}

func (wh *WriteHasher) Write(data []byte) (n int, err error) {
	return wh.w.Write(data)
}

func (wh *WriteHasher) Sum() []byte {
	return wh.h.Sum(nil)
}

func (wh *WriteHasher) Sum32() uint32 {
	return wh.h32.Sum32()
}

func (wh *WriteHasher) Sum64() uint64 {
	return wh.h64.Sum64()
}
