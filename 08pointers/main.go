package  main

import "fmt"

func main(){
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int
	// fmt.Println("Value of ptr is",ptr)
    myNum := 23
	var ptr = &myNum
	fmt.Println("Value of ptr is",ptr)
	fmt.Println("Value of myNum is",*ptr)
	fmt.Println("Value of &ptr is",&ptr)
	fmt.Println("value of *ptr is",*(&ptr))
   
	*ptr = *ptr *4
	fmt.Println("new value of myNum is",myNum)

	*ptr=*ptr+1
	fmt.Println("new value of myNum is",myNum)

}
