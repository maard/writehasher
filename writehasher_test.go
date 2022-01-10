package writehasher

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"
	"testing"
)

const data = "The quick brown fox jumps over the lazy dog"

func testHash(t *testing.T, name string, h hash.Hash, expected []byte) {
	var buf bytes.Buffer
	wh := NewWriteHasher(&buf, h)
	io.WriteString(wh, data)
	sum := wh.Sum()
	if !bytes.Equal(sum, expected) {
		t.Errorf("%s failed, got: %x, expected: %x\n", name, sum, expected)
	}
}

func testHash32(t *testing.T, name string, h hash.Hash32, expected uint32) {
	var buf bytes.Buffer
	wh := NewWriteHasher32(&buf, h)
	io.WriteString(wh, data)
	sum := wh.Sum32()
	if sum != expected {
		t.Errorf("%s failed, got: %d, expected: %d\n", name, sum, expected)
	}
}

func testHash64(t *testing.T, name string, h hash.Hash64, expected uint64) {
	var buf bytes.Buffer
	wh := NewWriteHasher64(&buf, h)
	io.WriteString(wh, data)
	sum := wh.Sum64()
	if sum != expected {
		t.Errorf("%s failed, got: %d, expected: %d\n", name, sum, expected)
	}
}

func TestAll(t *testing.T) {

	testHash(t, "crc IEEE", crc32.NewIEEE(), []byte("\x41\x4f\xa3\x39"))
	testHash(t, "crc64 ISO", crc64.New(crc64.MakeTable(crc64.ISO)), []byte("\x4e\xf1\x4e\x19\xf4\xc6\xe2\x8e"))
	testHash(t, "crc64 ECMA", crc64.New(crc64.MakeTable(crc64.ECMA)), []byte("\x5b\x5e\xb8\xc2\xe5\x4a\xa1\xc4"))
	testHash(t, "md5", md5.New(), []byte("\x9e\x10\x7d\x9d\x37\x2b\xb6\x82\x6b\xd8\x1d\x35\x42\xa4\x19\xd6"))
	testHash(t, "sha1", sha1.New(), []byte("\x2f\xd4\xe1\xc6\x7a\x2d\x28\xfc\xed\x84\x9e\xe1\xbb\x76\xe7\x39\x1b\x93\xeb\x12"))
	testHash(t, "sha256", sha256.New(), []byte("\xd7\xa8\xfb\xb3\x07\xd7\x80\x94\x69\xca\x9a\xbc\xb0\x08\x2e\x4f"+
		"\x8d\x56\x51\xe4\x6d\x3c\xdb\x76\x2d\x02\xd0\xbf\x37\xc9\xe5\x92"))
	testHash(t, "sha512", sha512.New(), []byte("\x07\xe5\x47\xd9\x58\x6f\x6a\x73\xf7\x3f\xba\xc0\x43\x5e\xd7\x69"+
		"\x51\x21\x8f\xb7\xd0\xc8\xd7\x88\xa3\x09\xd7\x85\x43\x6b\xbb\x64"+
		"\x2e\x93\xa2\x52\xa9\x54\xf2\x39\x12\x54\x7d\x1e\x8a\x3b\x5e\xd6"+
		"\xe1\xbf\xd7\x09\x78\x21\x23\x3f\xa0\x53\x8f\x3d\xb8\x54\xfe\xe6"))

	// 2fd4e1c67a2d28fced849ee1bb76e7391b93eb12
	// d7a8fbb307d7809469ca9abcb0082e4f8d5651e46d3cdb762d02d0bf37c9e592
	// 07e547d9586f6a73f73fbac0435ed76951218fb7d0c8d788a309d785436bbb642e93a252a954f23912547d1e8a3b5ed6e1bfd7097821233fa0538f3db854fee6

	testHash32(t, "crc IEEE", crc32.NewIEEE(), 0x414fa339)

	testHash64(t, "crc64 ISO", crc64.New(crc64.MakeTable(crc64.ISO)), 0x4ef14e19f4c6e28e)
	testHash64(t, "crc64 ECMA", crc64.New(crc64.MakeTable(crc64.ECMA)), 0x5b5eb8c2e54aa1c4)

}
