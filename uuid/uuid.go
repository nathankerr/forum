package uuid

import (
	"crypto/sha1"
	"encoding/hex"
)

var (
	nsUsers = []byte("1f24dbac-7f88-4937-bdb7-5641bd0831c8")
)

func NewV5(name string) string {
	hash := sha1.New()
	hash.Write(nsUsers[:])
	hash.Write([]byte(name))

	uuid := make([]byte, 16)
	buf := make([]byte, 36)

	copy(uuid[:], hash.Sum(nil))

	hex.Encode(buf[0:8], uuid[0:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])

	return string(buf)
}
