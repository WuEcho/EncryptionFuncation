package AES_CBC


import (
"bytes"
"crypto/aes"
"crypto/cipher"
)

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {

	padding := blocksize - len(ciphertext)%blocksize

	padtext := bytes.Repeat([]byte{byte(padding)},padding)

	return append(ciphertext,padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)

	unpadding := int(origData[length - 1])

	return origData[:(length-unpadding)]
}

//通过AES加密
func AESEncrypt(origData []byte, key []byte) []byte {
	//分组秘钥
	block,_ := aes.NewCipher(key)
	//获得秘钥块的长度
	blocksize := block.BlockSize()

	origData = PKCS7Padding(origData,blocksize)
	//加密模式
	blockMode := cipher.NewCBCEncrypter(block,key[:blocksize])
	crypted := make([]byte,len(origData))

	//加密
	blockMode.CryptBlocks(crypted,origData)

	return crypted
}

//AES解密
func AESDecrypt(cryed []byte,key []byte) []byte {
	block,_ := aes.NewCipher(key)

	blocksize := block.BlockSize()
	//设置解密方式
	blockMode := cipher.NewCBCDecrypter(block,key[:blocksize])

	origData := make([]byte,len(cryed))

	blockMode.CryptBlocks(origData,cryed)

	origData = PKCS7UnPadding(origData)

	return origData
}
