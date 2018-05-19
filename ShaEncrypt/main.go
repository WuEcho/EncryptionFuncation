package main

import (
	"EncryptDemo/ShaEncrypt/sha256Encrypt"
	"fmt"

)

func main() {

  var encryptCode =	sha256Encrypt.Sha256EncryptTypeOne([]byte("hello world"))

  fmt.Printf("%x",encryptCode)

  fmt.Println()

  var dencrypt = sha256Encrypt.Sha256EncryptTypeTwo([]byte("hello world"))

  fmt.Printf("%x",dencrypt)

  fmt.Println()
  fileEncrypted,err := sha256Encrypt.FileEncryptBySha256WithPath("/Users/wuxinyang/Desktop/MyGo/src/EncryptDemo/ShaEncrypt/test.txt")


  if err == nil {
		fmt.Println(fileEncrypted)
	}

}
