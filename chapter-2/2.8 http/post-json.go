package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Login struct {
	User     string
	Password string
}

func main() {

	login := Login{
		"sechelper",
		"123456",
	}

	loginJson, _ := json.Marshal(login)

	resp, _ := http.Post("http://go-hacker-code.lab.secself.com", "application/json",
		bytes.NewBuffer(loginJson))
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
}
