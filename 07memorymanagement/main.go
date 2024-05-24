package main
import "fmt"

func main() {
	// Creating a pointer
	//var p *int
	//p = new(int)
	//OR
	p := new(int) // Allocates memory for an int and returns a pointer

	 // Creating a slice with make
    s := make([]int, 0, 10) // Creates an empty slice with capacity 10
      fmt.Println(s)
    // Creating a map with make
    m := make(map[string]int) // Creates an empty map
	  fmt.Println(m ,"and",*p)


}
