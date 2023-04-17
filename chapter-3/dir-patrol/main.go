package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sechelper/seclib/async"
	"github.com/sechelper/seclib/dict"
	"github.com/sechelper/seclib/network"
	"log"
	"os"
)

func main() {
	path := "/Users/cookun/Downloads/dir.txt"
	op, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	d := dict.NewDict(op)

	async.Goroutine(10, func(c *chan any) {
		for d.Scan() {
			if line, err := d.Line(); err == nil {
				*c <- line
			}
		}
	}, func(args ...any) {
		url := "https://go-hacker-code.lab.secself.com/" + args[0].(string)

		resp, err := network.DefaultHttpClient.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode == 200 || resp.StatusCode == 302 ||
			resp.StatusCode == 301 {
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			title := doc.Find("title").Text()
			fmt.Println("[*]", resp.StatusCode, url, title)
		} else {
			fmt.Println("[x]", resp.StatusCode, url)
		}
		(*resp).Body.Close()
	})

}
