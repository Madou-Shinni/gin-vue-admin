package tools

import (
	"fmt"
	"testing"
)

func TestSha1Encrypt(t *testing.T) {
	inputString := "string1"

	// 对字符串进行 SHA-1 加密
	encryptedString := Sha1Encrypt(inputString)

	fmt.Printf("Original String: %s\n", inputString)
	fmt.Printf("SHA-1 Encrypted String: %s\n", encryptedString)
}
