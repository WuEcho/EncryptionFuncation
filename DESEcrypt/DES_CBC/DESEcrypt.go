package DES_CBC

import (
	"crypto/cipher"
	"crypto/des"
	"bytes"
)

//DES加密方法
func DesEncrypt(origData, key []byte) []byte{

	block, _ := des.NewCipher(key)

	//将明文按秘钥的长度做补码运算

	origData = PKCS5Padding(origData, block.BlockSize())

	//设置加密方式
	//NewCBCEncrypter
	blockMode := cipher.NewCBCDecrypter(block, key)

	//创建明文长度的字节数组

	crypted := make([]byte, len(origData))

	//加密明文

	blockMode.CryptBlocks(crypted, origData)

	return crypted
}

//实现明文的补码

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {

	padding := blockSize - len(ciphertext)%blockSize

	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)

}


//DES解密方法
func DESDecrypt(crypted []byte, key []byte) []byte {


	//将字节秘钥转换成block快
	block, _ := des.NewCipher(key)

	//设置解密方式
	blockMode := cipher.NewCBCEncrypter(block, key)

	//创建密文大小的数组变量
	origData := make([]byte, len(crypted))

	//解密密文到数组origData中
	blockMode.CryptBlocks(origData, crypted)

	//去补码
	origData = PKCS5UnPadding(origData)

    return origData
}

//实现去补码

func PKCS5UnPadding(origData []byte) []byte {

	length := len(origData)

	// 去掉最后一个字节 unpadding 次

	unpadding := int(origData[length-1])

	return origData[:(length - unpadding)]
}


