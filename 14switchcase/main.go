package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang")

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's a weekday")
	// }

	// // switch without a condition is an alternate way of switch-case

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("Good morning!")
	// case t.Hour() < 17:
	// 	fmt.Println("Good afternoon!")
	// default:
	// 	fmt.Println("Good evening!")
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a bool")
	// 	case int:
	// 		fmt.Println("I'm an int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }
	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

	rand.Seed(time.Now().UnixNano())
	dicenum := rand.Intn(6) + 1
	fmt.Println("dice number is ", dicenum)

	switch dicenum {
	case 1:
		fmt.Println("you got 1 and you can move 1 spot or open")
	case 2:
		fmt.Println("you got 2 and you can move 2 spot")
	case 3:
		fmt.Println("you got 3 and you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("you got 4 and you can move 4 spot")
		fallthrough
	case 5:
		fmt.Println("you got 5 and you can move 5 spot")
		fallthrough
	case 6:
		fmt.Println("you got 6 and you can move 6 spot and you can roll the dice again")
	}

}
