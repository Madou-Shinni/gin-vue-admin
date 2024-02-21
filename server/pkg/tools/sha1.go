package tools

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1Encrypt(input string) string {
	// 创建 SHA-1 hasher
	hasher := sha1.New()

	// 将字符串转换为字节数组并写入 hasher
	hasher.Write([]byte(input))

	// 获取 SHA-1 散列值
	hashBytes := hasher.Sum(nil)

	// 将散列值转换为十六进制字符串
	hashString := hex.EncodeToString(hashBytes)

	return hashString
}
