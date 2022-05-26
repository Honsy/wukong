package lib

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))

	return hex.EncodeToString(m.Sum(nil))
}
