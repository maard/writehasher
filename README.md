# writehasher

A drop-in replacement to any Writer type, which also calculates a hash using the provided hash type.

# Example

```go
package main

import (
	"fmt"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"

	"github.com/maard/writehasher"
)

func main() {
	var w io.Writer = io.Discard
	data := "the hash of this data needs to be calculated"

	// Usage with types, which implement hash.Hash, for instance:
	// hasher := NewWriteHasher(w, crc32.NewIEEE())
	// hasher := NewWriteHasher(w, crc64.New(crc64.MakeTable(crc64.ISO)))
	// hasher := NewWriteHasher(w, crc64.New(crc64.MakeTable(crc64.ECMA)))
	// hasher := NewWriteHasher(w, md5.New())
	// hasher := NewWriteHasher(w, sha1.New())
	// hasher := NewWriteHasher(w, sha256.New())
	// hasher := NewWriteHasher(w, sha512.New())

	hasher := NewWriteHasher(w, md5.New())
	io.WriteString(hasher, data)
	fmt.Printf("%x\n", hasher.Sum())

	hasher = NewWriteHasher(w, crc32.NewIEEE())
	io.WriteString(hasher, data)
	fmt.Printf("%x\n", hasher.Sum())

	// Usage with types, which also implement hash.Hash32, for instance:
	hasher = NewWriteHasher32(w, crc32.NewIEEE())
	io.WriteString(hasher, data)
	fmt.Printf("%x\n", hasher.Sum32())

	// Usage with types, which also implement hash.Hash64, for instance:
	hasher = NewWriteHasher64(w, crc64.New(crc64.MakeTable(crc64.ISO)))
	io.WriteString(hasher, data)
	fmt.Printf("%x\n", hasher.Sum64())
}
// prints:
// b9a9aa9d96160f293562ffcfc543bc1d
// 8dfa06fe
// 8dfa06fe
// 7fac11fe2939f29
```

# API

### hasher := NewWriteHasher(w io.Writer, h hash.Hash)
### hasher := NewWriteHasher32(w io.Writer, h hash.Hash32)
### hasher := NewWriteHasher64(w io.Writer, h hash.Hash64)

Creates a class instance.

### hasher.Sum()
### hasher.Sum32()
### hasher.Sum64()

Returns the hash of the data, written to hasher (`[]byte`, `uint32`, `uint64` respectively)
