package AES_CFB

//CFB模式进行AES加密

import (

"crypto/aes"
"io"
"crypto/rand"
"crypto/cipher"

)

//加密
func AESEcrypt(plaintext []byte, key []byte) []byte {
	//分组秘钥
	block, _:= aes.NewCipher(key)
	//创建了个数组 用来存放加密后的密文
	ciphertext := make([]byte,aes.BlockSize+len(plaintext))
	//设置内存空间可读
	iv := ciphertext[:aes.BlockSize]
	io.ReadFull(rand.Reader,iv)
	//设置加密模式 会返回一个流
	stream := cipher.NewCFBEncrypter(block,iv)
	//加密利用ciphertext[aes.BlockSize:]与plaintext进行异或运算
	stream.XORKeyStream(ciphertext[aes.BlockSize:],plaintext)

	return ciphertext
}

//解密
func AESDecrypt(ciphertext []byte,key []byte) []byte {

	block,_:= aes.NewCipher(key)

	iv := ciphertext[:aes.BlockSize]

	ciphertext = ciphertext[aes.BlockSize:]
	//设置解密方式
	stream := cipher.NewCFBDecrypter(block,iv)

	stream.XORKeyStream(ciphertext,ciphertext)

	return ciphertext
}


