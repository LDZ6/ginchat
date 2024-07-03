package utils

import (
	"crypto/md5"
	"strings"
)

// 返回值为小写的MD5值
func Md5Encode(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	tempStr := h.Sum(nil)
	md5Str := string(tempStr)
	return md5Str
}

// 返回值为大写的MD5值
func MD5Encode(str string) string {
	return strings.ToUpper(Md5Encode(str))
}

// 生成密码
func MakePassword(plainpwd, salt string) string {
	return MD5Encode(plainpwd + salt)
}

// 解密
func ValidPassword(encryptedpwd, salt string, passwoed string) bool {
	return MD5Encode(encryptedpwd+salt) == passwoed
}
