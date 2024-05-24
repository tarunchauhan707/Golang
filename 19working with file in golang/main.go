package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//reading or writing files in golang

	//open file
	//create file

	content := "This needs to go in a file"
	file, err := os.Create("./MYGOFILE.txt")

	if err != nil {
		panic(err)
	}
	//write to file
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is", length)

	//close file
	defer file.Close()
	readFile("./MYGOFILE.txt")
}

// read file in golang
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is \n", string(databyte))
}
