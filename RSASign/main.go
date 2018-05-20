package main

import (
	"crypto/rsa"
	"crypto/md5"
	"crypto/rand"
	"crypto"
	"fmt"
)

func main() {

	//生成私钥
	priv ,_ := rsa.GenerateKey(rand.Reader,1024)
	//生产公钥
	pub := &priv.PublicKey

	//设置明文
	plaintText := []byte("hello world")

	h := md5.New()
	h.Write(plaintText)

	hashed := h.Sum(nil)

	//签名
	opts := &rsa.PSSOptions{SaltLength:rsa.PSSSaltLengthAuto,Hash:crypto.MD5}
	sing, _ := rsa.SignPSS(rand.Reader,priv,crypto.MD5,hashed,opts)

	//认证
	e := rsa.VerifyPSS(pub,crypto.MD5,hashed,sing,opts)

	if e == nil {
		fmt.Println("验证成功")
	}
}
