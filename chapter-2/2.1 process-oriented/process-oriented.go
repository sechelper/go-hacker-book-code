package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

//func MD5(str string) string {
//	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
//}

func SHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	// 函数作用是将十六进制bs变量转换成字符串
	return fmt.Sprintf("%x", bs)

}

func SHA512(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	// 函数作用是将十六进制bs变量转换成字符串
	return fmt.Sprintf("%x", bs)
}

func main() {
	//fmt.Println(MD5("123456"))
	fmt.Println(SHA256("123456"))
	fmt.Println(SHA512("123456"))
}
