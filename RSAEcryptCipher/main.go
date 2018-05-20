package main

import (
	"EncryptDemo/RSAEcryptCipher/RSACipher"
	"fmt"
	"encoding/base64"
)


func main() {

	data,_ := RSACipher.RsaEncrypt([]byte("hello world"))
	fmt.Println(base64.StdEncoding.EncodeToString(data))

	origData,_ := RSACipher.RsaDecrypt(data)
	fmt.Println(string(origData))
}
