package RSACipher

import (
	"crypto/rsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
)

var privateKey *rsa.PrivateKey = new(rsa.PrivateKey)

var publishKey *rsa.PublicKey = new(rsa.PublicKey)


//加密
func RSAEcrypt(origData []byte) []byte {

	//生成私钥
	privateKey,_ = rsa.GenerateKey(rand.Reader,1024)

	//生成公钥
	publishKey = &privateKey.PublicKey

	//加密
	cipherText,_ := rsa.EncryptOAEP(md5.New(),rand.Reader,publishKey,origData,nil)

	return cipherText
}


//解密
func RSADecrypt(origData []byte) []byte {

	plainText,err:= rsa.DecryptOAEP(md5.New(),rand.Reader,privateKey,origData,nil)
    if err != nil{
    	fmt.Println(err)

	}

	return plainText
}