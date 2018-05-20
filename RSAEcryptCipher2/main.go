package main

import (
	"fmt"
	"EncryptDemo/RSAEcryptCipher2/RSACipher"
	"encoding/base64"
)
func main() {

	encryptCode := RSACipher.RSAEcrypt([]byte("hello world"))

	fmt.Println(base64.StdEncoding.EncodeToString(encryptCode))

	dencryptCode := RSACipher.RSADecrypt(encryptCode)

	fmt.Println(string(dencryptCode))
}
