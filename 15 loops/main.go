package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	//there is no ++i or --i , it is always i++ or i--
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index, day)
	// }
	//comma ok syntax
	// for _, day := range days {
	// 	fmt.Printf("value is %v \n", day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			break
		}
		if rougueValue == 8 {
			rougueValue++ //alone continue is not working here
			continue
		}
		fmt.Println("Value is ", rougueValue)
		rougueValue++
	}
lco:
	fmt.Println("Jumping at lco")
}
