package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	// no inheritance in golang
	// no private variables
	// no private methods
	// no overloading
	// no operator overloading
	//No super or parent class
	//No multiple inheritance
	dev := User{"dev", "dev@go.dev", 23, "1234567890", true}
	fmt.Println(dev)
	fmt.Println("details of dev")
	fmt.Println(" Name is", dev.Name,"\n","Email is",dev.Email,"\n", "Age is ", dev.Age,"\n", "Phone is ", dev.Phone,"\n", "Status is", dev.Status)

}

type User struct {
	Name   string
	Email  string
	Age    int
	Phone  string
	Status bool
}
