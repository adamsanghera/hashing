package hashing

import (
	"encoding/hex"
)

//WithNewSalt ...
// returns ENCODED salt and ENCODED hash
func WithNewSalt(pass string) (string, string) {
	salt := generateSalt()

	hash := runScryptWithSalt(pass, salt)

	return hex.EncodeToString(salt), hex.EncodeToString(hash)
}

//WithOldSalt ...
// expects password, ENCODED salt
// returns ENCODED hash
func WithOldSalt(pass string, encodedSalt string) string {
	salt := decode(encodedSalt)

	hash := runScryptWithSalt(pass, salt)

	return hex.EncodeToString(hash)
}

//IsValidChallenge ...
//  expects a challenge, encoded salt, encoded hash.
func IsValidChallenge(challenge string, encodedSalt string, encodedHash string) bool {
	return WithOldSalt(challenge, encodedSalt) == encodedHash
}

// WithoutSalt is used to hash strings without salt.
// This isn't really recommended.
// If you have the capacity to store salt, you should use WithNewSalt instead.
func WithoutSalt(pass string) string {
	return hex.EncodeToString(runScryptWithoutSalt(pass))
}
