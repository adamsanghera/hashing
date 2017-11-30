package hashing

import (
	"encoding/hex"

	"golang.org/x/crypto/scrypt"
)

func runScrypt(pass string, salt []byte) []byte {
	hash, err := scrypt.Key([]byte(pass), salt, 1<<14, 8, 1, pwHashBytes)
	if err != nil {
		panic(err)
	}
	return hash
}

func decode(encodedStr string) []byte {
	decodedStr := make([]byte, hex.DecodedLen(len(encodedStr)))
	_, err := hex.Decode(decodedStr, []byte(encodedStr))
	if err != nil {
		panic(err)
	}

	return decodedStr
}
