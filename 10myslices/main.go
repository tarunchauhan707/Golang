package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

    var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("type of numbers is %T\n ", numbers)  //here println do not return the type of the variable

	numbers = append(numbers, 11, 12, 13, 14, 15)
	fmt.Println("numbers are: ", numbers)

	numbers=append(numbers[:3], numbers[8:]...)
	fmt.Println("numbers are: ", numbers)

	num:=make([]int, 6)
	num[0]=124
	num[1]=234
	num[2]=465
	num[3]=432
	num[4]=566
	num[5]=234
	//num[6]=234
	fmt.Println(num)
	// fmt.Println(sort.IntsAreSorted(num))

	num=append(num, 707, 808, 909)
	fmt.Println(num)
	//sorting the slice 
    sort.Ints(num)
	fmt.Println(num)
	fmt.Println(sort.IntsAreSorted(num))

	//how to remove an element from slice

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice)
     // 1st way using append method
	var index int =2
	slice=append(slice[:index], slice[index+1:]...)
	fmt.Println(slice)
    

}