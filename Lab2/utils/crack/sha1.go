package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(password string) string {
	digest := sha1.Sum([]byte(password))
	return hex.EncodeToString(digest[:])
}
