package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
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
