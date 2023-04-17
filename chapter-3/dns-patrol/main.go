package main

import (
	"fmt"
	"github.com/sechelper/seclib/async"
	"github.com/sechelper/seclib/dict"
	"github.com/sechelper/seclib/network"
	"log"
	"os"
	"time"
)

func main() {
	domain := "secself.com"

	resolver := network.Dns{
		NewMsg:  network.NewDefaultMsg,
		Ns:      "8.8.8.8",
		Timeout: 5 * time.Second,
	}

	async.Goroutine(10, func(c *chan any) {
		path := "dns.txt"
		op, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		dt := dict.NewDict(op)
		if err != nil {
			log.Fatal(err)
		}

		for dt.Scan() {
			*c <- dt.Text()
		}
	}, func(a ...any) {

		ips, err := resolver.LookupIP(a[0].(string) + "." + domain)
		if err != nil {
			return
		}
		if len(ips) != 0 {
			fmt.Println("[*]", a[0].(string)+"."+domain, ips)
		} else {
			fmt.Println("[x]", a[0].(string)+"."+domain)
		}
	})

}
