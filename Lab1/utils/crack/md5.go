package crack

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(password string) string {
	digest := md5.Sum([]byte(password))
	return hex.EncodeToString(digest[:])
}
