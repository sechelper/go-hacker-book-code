package api

import (
	"fmt"
	scope2 "go-scope/scope"
)

func Auth() {
	fmt.Println(scope2.Username)
	scope2.Login("admin", "admin")
}
