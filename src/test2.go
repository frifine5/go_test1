package main

import (
	"fmt"
	"encoding/base64"
	"log"
)

func main() {
	fmt.Println("adfad")
	s1 := string("dfada")
	fmt.Println(s1 + ":00000000")
	i1 :=int8(0x7f)
	fmt.Println(i1)
	i2 := int16(0x7fff)
	fmt.Println(i2)
	fmt.Println(uint8(0xff))
	fmt.Println(uint16(0xffff))
	fmt.Println(uint32(0xffffffff))
	fmt.Println(uint64(0xffffffffffffffff))
	fmt.Println("---------------------->")

	base64test()
}

func base64test(){
	input := []byte("hello world=+?&%")

	// 演示base64编码
	fmt.Println("使用StdEncoding演示base64编码>>>>>")
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Printf( "编码：%s\n", encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf( "解码：%s\n", string(decodeBytes))


	// 如果要用在url中，需要使用URLEncoding
	fmt.Println("使用URLEncoding演示base64编码>>>>>")
	uEnc := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Printf( "编码：%s\n", uEnc)

	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("解码：%s\n", string(uDec))
}
