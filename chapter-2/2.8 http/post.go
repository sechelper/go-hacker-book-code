package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {

	data := url.Values{
		"user": {"sechelper"},
	}

	resp, err := http.PostForm("http://go-hacker-code.lab.secself.com", data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
}
