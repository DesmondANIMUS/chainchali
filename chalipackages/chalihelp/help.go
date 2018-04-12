package chalihelp

import (
	"crypto/sha256"
	"encoding/hex"
)

// GetHash function returns hash for the current block
func GetHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
