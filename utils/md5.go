package utils

import (
	md52 "crypto/md5"
	"encoding/hex"
)

func MD5(src []byte) string {
	m := md52.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
