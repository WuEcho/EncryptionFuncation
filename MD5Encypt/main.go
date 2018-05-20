package main

import (
	"EncryptDemo/MD5Encypt/MD5"
	"fmt"
)

func main() {

	var encryptCode = MD5.MD5EncryptTypeOne([]byte("hello world"))

	fmt.Println(encryptCode)

	var encryptCode1 = MD5.MD5EncryptTypeTwo([]byte("hello world"))

	fmt.Println(encryptCode1)

}
