package main

import (
	"crypto/rand"
	"fmt"
	"crypto/dsa"
)

func main() {

	//DSA专业签名验签
	var param dsa.Parameters

	//配置
	dsa.GenerateParameters(&param,rand.Reader,dsa.L1024N160)

	//生成私钥
	var priv dsa.PrivateKey
	priv.Parameters = param
	dsa.GenerateKey(&priv,rand.Reader)

	//通过私钥生公钥
	pub := priv.PublicKey

	//利用私钥签名数据
	message := []byte("helloworld")
	r,s,_ := dsa.Sign(rand.Reader,&priv,message)

	//公钥验签
	b := dsa.Verify(&pub,message,r,s)

	if b == true {
		fmt.Println("验签成功")
	}

}
