package main

import (
	"EncryptDemo/Ripemd160Ecrypt/Ripemd160"
	"fmt"
)

//ripemd160  hash加密
//需要导入三方包
//打开终端
//cd $GOPATH/src
//cd golang.org
//cd x
//git clone https://github.com/golang/crypto.git

func main() {

	var ripemdEcrypt = Ripemd160.Ripemd160Ecrypt([]byte("hello world"))

	fmt.Println(ripemdEcrypt)
}
