package main
import "fmt"
import "bufio"
import "os"

func main() {
	welcome:= "user input from user"
	fmt.Println("hello, ok, comma, syntax")
	fmt.Println(welcome)

    //taking user input
	reader:= bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value:")
    
    //comma ok syntax || error ok syntax
	text, _ := reader.ReadString('\n')         	//reading the user input as a string untill '\n'
	fmt.Println("Thanks for entering : ", text)
}