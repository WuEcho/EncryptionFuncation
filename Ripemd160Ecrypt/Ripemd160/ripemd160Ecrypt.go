package Ripemd160

import (
	"golang.org/x/crypto/ripemd160"
	"fmt"
)

func Ripemd160Ecrypt(origData []byte) string {
	hasher :=  ripemd160.New()
	hasher.Write(origData)

	hashString := fmt.Sprintf("%x",hasher.Sum(nil))

	return hashString
}
