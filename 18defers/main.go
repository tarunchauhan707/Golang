package main

import "fmt"

func main() {
	//any statement that is preceded by the
	//defer keyword isn’t invoked until the
	// end of the function in which defer was used.
	defer fmt.Println("-------------------------------")
	defer fmt.Println("World")

	fmt.Println("Hello")

	//defer executes in LIFO order

	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")

	fmt.Println("Hello")
	mydefer()

	//hello , hello ,mydefetr, three, two, one, world,-------------------

}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
