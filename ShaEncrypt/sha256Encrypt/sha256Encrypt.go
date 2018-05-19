package sha256Encrypt

import (
	"crypto/sha256"
	"os"
	"fmt"
	"io"
)

func Sha256EncryptTypeOne(origData []byte) [sha256.Size]byte {

	sum := sha256.Sum256(origData)

	return sum
}

func Sha256EncryptTypeTwo(origData []byte) []byte {

	h := sha256.New()

	h.Write(origData)

	return h.Sum(nil)
}

func FileEncryptBySha256WithPath(filePath string) (string,error){

	//文件路径必须为绝对路径
	f,err := os.Open(filePath)

	if err != nil {
		fmt.Print(err)
		return "",err
	}

	defer f.Close()

	h := sha256.New()

	io.Copy(h,f)

	s := fmt.Sprintf("%x",h.Sum(nil))

	return s,err
}