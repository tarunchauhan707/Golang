package main

import "fmt"

//declaring const -
//first letter is capital->public
const LoginToken string = "abc123"

func main(){


	var username string="Golang"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var smallnum uint8 = 255
	fmt.Println(smallnum)
	fmt.Printf("Variable is of type: %T \n",smallnum)


	var isask bool = true
	fmt.Println(isask)
	fmt.Printf("Variable is of type: %T \n",isask)

	var smallflaot float32 = 255.0
	fmt.Println(smallflaot)
	fmt.Printf("Variable is of type: %T \n",smallflaot)

	var bigflaot float64 = 255.0
	fmt.Println(bigflaot)
	fmt.Printf("Variable is of type: %T \n",bigflaot)

	//// implicit type
	var website = "google.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n",website)



	//no var use
	numberofuser:= 300
	fmt.Println(numberofuser)


	fmt.Println(LoginToken)


}

