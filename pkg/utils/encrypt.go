package utils

import (
	"crypto/md5"
	"mygo/config"
)

func Encrypt(password string) string {
	salt := config.Server.Salt
	hash := md5.Sum([]byte(salt + password))
	return string(hash[:])
}

func CmpPwd(hash string, password string) bool {
	return hash == Encrypt(password)
}
