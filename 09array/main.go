package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	//declaring of array
	var arrlist [5]string

	arrlist[0] = "Golang"
	arrlist[1] = "Java"
	arrlist[3] = "Python"
	arrlist[4] = "Ruby"

	fmt.Println("all array values are: ",arrlist)
    fmt.Println("array length is: ",len(arrlist))


	arrlist[2] = "C++"

	fmt.Println("all array values are: ",arrlist)
    fmt.Println("array length is: ",len(arrlist))

	var myarray = [5]int{1,2,3,5}
	fmt.Println("all array values are: ",myarray)
	fmt.Println("array length is: ",len(myarray))
}