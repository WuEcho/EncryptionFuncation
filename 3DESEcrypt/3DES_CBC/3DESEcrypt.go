package _DES_CBC

import (
"bytes"
"crypto/des"
"crypto/cipher"
)

func PKCS5Padding(ciphertext []byte,blocksize int) []byte {

	padding := blocksize - len(ciphertext)%blocksize

	padtext := bytes.Repeat([]byte{byte(padding)},padding)

	return append(ciphertext,padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {

	length := len(origData)

	unadding := int(origData[length-1])

	return origData[:(length - unadding)]
}

func TripleDesEncrypt(origData []byte,key []byte) []byte  {

	//秘钥长度必须是24位
	block, _ := des.NewTripleDESCipher(key)

	//补码
	origData = PKCS5Padding(origData,block.BlockSize())

	//设置加密模式
	blockMode := cipher.NewCBCEncrypter(block,key[:8])

	//创建密文数组，加密
	crypted := make([]byte,len(origData))

	//加密
	blockMode.CryptBlocks(crypted,origData)

	return crypted
}

//解密
func TripleDesDecrypt(crypted []byte, key []byte) []byte{
	//秘钥长度必须是24位
	block, _ := des.NewTripleDESCipher(key)

	//设置加密模式
	blockMode := cipher.NewCBCDecrypter(block,key[:8])

	//创建密文数组，加密
	origData := make([]byte,len(crypted))

	//加密
	blockMode.CryptBlocks(origData,crypted)

	origData = PKCS5UnPadding(origData)

	return origData
}

