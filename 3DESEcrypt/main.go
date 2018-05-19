package main

import (
	"fmt"
	"EncryptDemo/3DESEcrypt/3DES_CBC"
	"encoding/base64"
)

func main()  {

	var key []byte = []byte("123456789012345678901234")

	var encryptCode = _DES_CBC.TripleDesEncrypt([]byte("helloworld"),key)

	fmt.Println(base64.StdEncoding.EncodeToString(encryptCode))

	var decryptCode =  _DES_CBC.TripleDesDecrypt(encryptCode,key)

	fmt.Println(string(decryptCode))
}
