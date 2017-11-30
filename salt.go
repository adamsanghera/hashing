package hashing

import (
	"crypto/rand"
	"io"
)

const (
	pwSaltBytes = 32
	pwHashBytes = 256
)

func generateSalt() []byte {
	salt := make([]byte, pwSaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		panic(err)
	}
	return salt
}
