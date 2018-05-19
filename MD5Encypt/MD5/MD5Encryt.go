package MD5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func MD5EncryptTypeOne(origeData []byte) string {

	h := md5.New()
	h.Write(origeData)
	//16进制转字符串
	s := hex.EncodeToString(h.Sum(nil))

   return s
}

func MD5EncryptTypeTwo(origeData []byte) string {

	s := fmt.Sprintf("%x",md5.Sum(origeData))

	return s
}

