package main

import (
	"fmt"
    "EncryptDemo/AESCBCEcrypt/AES_CBC"
	"encoding/base64"
)

func main() {

		//16
		//var key =[]byte("1234567890123456")

		//24
		var key2 = []byte("123456789012345678901234")

		//32
		//var key1 = []byte("12345678901234567890123456789012")

		var encryCode = AES_CBC.AESEncrypt([]byte("HELLO"),key2)

		fmt.Println(base64.StdEncoding.EncodeToString(encryCode))

		var dencryptCode = AES_CBC.AESDecrypt(encryCode,key2)

		fmt.Println(string(dencryptCode))
}