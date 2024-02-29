package utils

import (
	"crypto/md5"
	"encoding/hex"
	"mygo/config"
)

func Encrypt(password string) string {
	salt := config.Server.Salt
	hash := md5.Sum([]byte(salt + password))
	return hex.EncodeToString(hash[:])
}

func CmpPwd(hash string, password string) bool {
	return hash == Encrypt(password)
}
