package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	// encode
	hello := "你好，世界！ hello world"
	enByte := base64Encode([]byte(hello))
	fmt.Println(enByte, string(enByte))
	// decode
	deByte, err := base64Decode(enByte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(deByte) {
		fmt.Println("hello is not equal to deByte")
	}

	fmt.Println(string(deByte))
}
