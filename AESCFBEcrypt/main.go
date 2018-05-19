package main

import (
	"fmt"
	"EncryptDemo/AESCFBEcrypt/AES_CFB"
	"encoding/base64"
)

func main()  {


	key := []byte("123456789abcdefg")

	var encryptCode = AES_CFB.AESEcrypt([]byte("hahha"),key)

	fmt.Println(base64.StdEncoding.EncodeToString(encryptCode))

	var decryptCode = AES_CFB.AESDecrypt(encryptCode,key)

	fmt.Println(string(decryptCode))
}
