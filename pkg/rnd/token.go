// Package rnd generates random strings good for use in URIs to identify
// unique objects.
//
// Example usage:
//
//	s := rnd.New() // s is now "apHCJBl7L1OmC57n"
//
// A standard string created by New() is 16 bytes in length and consists of
// Latin upper and lowercase letters, and numbers (from the set of 62 allowed
// characters), which means that it has ~95 bits of entropy. To get more
// entropy, you can use NewLen(UUIDLen), which returns 20-byte string, giving
// ~119 bits of entropy, or any other desired length.
//
// Functions read from crypto/rand random source, and panic if they fail to
// read from it.

package rnd

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
)

const (
	// StdLen is a standard length of rnd string to archive ~95 bits of entropy.
	StdLen = 16
	// UUIDLen is a length of rnd string to archive ~119 bits of entropy, closest
	// to what can be losslessly converted to UUIDv4 (122 bits).
	UUIDLen = 20
)

// StdChars is a set of standard characters allowed in rnd string.
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

var AsciiLowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var AsciiUppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var AsciiLetters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var Digits = []byte("0123456789")
var HexDigits = []byte("0123456789abcdefABCDEF")

// New returns a new random string of the standard length, consisting of
// standard characters.
func New() string {
	return NewLenChars(StdLen, StdChars)
}

// NewLen returns a new random string of the provided length, consisting of
// standard characters.
func NewLen(length uint) string {
	return NewLenChars(length, StdChars)
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters (maximum 256).
func NewLenChars(length uint, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("rnd: wrong charset length for NewLenChars")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	var i uint
	for {
		if _, err := rand.Read(r); err != nil {
			panic("rnd: error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

func Token(size uint) string {
	if size > 10 || size < 1 {
		panic(fmt.Sprintf("size out of range: %d", size))
	}

	result := make([]byte, 0, 14)
	b := make([]byte, 8)

	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	randomInt := binary.BigEndian.Uint64(b)

	result = append(result, strconv.FormatUint(randomInt, 36)...)

	for i := len(result); i < cap(result); i++ {
		result = append(result, byte(123-(cap(result)-i)))
	}

	return string(result[:size])

}
