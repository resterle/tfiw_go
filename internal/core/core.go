package core

import (
	"crypto/rand"
	"encoding/base32"
)

func RandomId() string {
	b := make([]byte, 5)
	rand.Read(b)
	return base32.HexEncoding.EncodeToString(b)
}
