package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func GetHash(arr []string) string {
	s := strings.Join(arr, "")
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}