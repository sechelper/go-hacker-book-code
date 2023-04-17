package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

var a string = "test"

type Hash struct {
	Str   string
	Count int
}

// MD5 e10adc3949ba59abbe56e057f20f883e
func (hash *Hash) MD5() string {
	hash.Count++
	hash.Str = "12345678" // 25d55ad283aa400af464c76d713c07ad
	return fmt.Sprintf("%x", md5.Sum([]byte(hash.Str)))
}

// SHA256 8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92
func (hash *Hash) SHA256() string {
	hash.Count++
	h := sha256.New()
	h.Write([]byte(hash.Str))
	bs := h.Sum(nil)
	// 函数作用是将十六进制bs变量转换成字符串
	return fmt.Sprintf("%x", bs)

}

func (hash *Hash) SHA512() string {
	hash.Count++
	h := sha512.New()
	h.Write([]byte(hash.Str))
	bs := h.Sum(nil)
	// 函数作用是将十六进制bs变量转换成字符串
	return fmt.Sprintf("%x", bs)
}

func main() {
	hash := Hash{Str: "123456"}
	fmt.Println(hash.Str)
	fmt.Println(hash.Count)
	fmt.Println(hash.MD5())
	fmt.Println(hash.SHA256())
	fmt.Println(hash.SHA512())
	fmt.Println(hash.Count)
}
