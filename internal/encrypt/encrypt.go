package encrypt

import (
	"fiber-api/config"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// 加密
func Generate(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password + signature()))
}

// 比较
func Compare(hashPassword, password string) bool {
	return base64.StdEncoding.EncodeToString([]byte(password+signature())) == hashPassword
}

// 签名
func signature() string {

	appKey := md5.Sum([]byte(config.App.Key))

	return hex.EncodeToString(appKey[:])
}
