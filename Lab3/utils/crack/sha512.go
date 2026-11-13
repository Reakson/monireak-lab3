package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

func Hash(password string) string {
	digest := sha512.Sum512([]byte(password))
	return hex.EncodeToString(digest[:])
}
