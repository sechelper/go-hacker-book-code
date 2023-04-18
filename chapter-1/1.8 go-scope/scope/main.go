package scope2

import "fmt"

var Username = "admin"

var password = "admin"
var password2 = "admin2"

func Login(username string, password string) {
	var realPwd = "admin"
	var password2 = "admin"
	fmt.Println(Username, username, password, realPwd, password2)
}

func logout() {
	fmt.Println("logout")
	fmt.Println(password)
}
