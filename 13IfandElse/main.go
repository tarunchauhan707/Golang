package main

import "fmt"

func main() {
	fmt.Println("If and Else")
    loginCount := 23
	var result string
	if (loginCount < 10) {
		result = "Regular user"
	} else if (loginCount > 10) {
		result = "Watch out"
	} else {
		result = "Exactly 10 logins"
	}
	fmt.Println(result)
}
