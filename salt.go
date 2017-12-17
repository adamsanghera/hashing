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

// GetSaltSize returns the size of the salt.
// Useful if you are combining hash and salt in common string.
func GetSaltSize() int {
	return pwSaltBytes * 4
}

// GetHashSize returns the size of the resulting hashed string.
// Useful if you are combining hash and salt in common string.
func GetHashSize() int {
	return pwHashBytes * 4
}
