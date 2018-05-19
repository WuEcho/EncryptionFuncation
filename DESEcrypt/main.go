package main

import (
	"fmt"
	"EncryptDemo/DESEcrypt/DES_CBC"
	"encoding/base64"
)
func main() {

	fmt.Println("Hello World!")

	//声明秘钥,利用此秘钥实现明文的加密和密文的解密

	key := []byte("12345678")

	//加密

	var encryptCode = DES_CBC.DesEncrypt([]byte("hello world"), key)

	fmt.Println(base64.StdEncoding.EncodeToString(encryptCode))

	//解密

	var decrypt = DES_CBC.DESDecrypt(encryptCode, key)

	fmt.Println(string(decrypt))
}