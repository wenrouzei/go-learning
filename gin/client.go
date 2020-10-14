package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	char := '你'
	v1 := rune(char)
	v2 := byte(char)
	s1 := strconv.FormatInt(int64(v1), 2)
	s2 := strconv.FormatInt(int64(v2), 2)
	fmt.Printf("v1: %c, type: %T, %v\n", v1, v1, s1)
	fmt.Printf("v2: %c, type: %T, %v\n", v2, v2, s2)
	fmt.Println(byte('a'))
	b := bytes.NewBuffer([]byte{})
	type a struct {
		B bool `json:"b"`
	}
	c := a{}
	c.B = true
	json.NewEncoder(b).Encode(c)

	request, _ := http.NewRequest("GET", "http://localhost:8080/", b)
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	rep, _ := client.Do(request)
	defer rep.Body.Close()
	body, _ := ioutil.ReadAll(rep.Body)
	fmt.Println(string(body))

	d, _ := json.Marshal(c)
	request2, _ := http.NewRequest("GET", "http://localhost:8080/", bytes.NewReader(d))
	request2.Header.Set("Content-Type", "application/json")
	rep2, _ := client.Do(request2)
	defer rep2.Body.Close()
	body2, _ := ioutil.ReadAll(rep2.Body)
	fmt.Println(string(body2))
}
