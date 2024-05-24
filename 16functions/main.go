package main

import "fmt"

func main() {
	fmt.Println("welcome to functions in golang")
	greeting()
	//-----you can not define a function inside another function----
	

	result := getSum(3, 4)
	fmt.Println("result is", result)


	proResult := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("proResult is", proResult)

	proResult2,_ := proAdder2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("proResult is", proResult2)
	//0R

	proResult3,mystring := proAdder2(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("proResult is", proResult3)
	fmt.Println("mystring is", mystring)
}

func getSum(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
func proAdder(values ...int) int {
	sum := 0	
	for _, value := range values {
		sum += value
	}
	return sum
}
//two types of values return function
func proAdder2(values ...int) (int, string) {
	sum := 0	
	for _, value := range values {
		sum += value
	}
	return sum,"i am returning sum"

}

func greeting() {
	fmt.Println("hello there")
}
