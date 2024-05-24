package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//---net/http package

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to handling web requests in golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//
	fmt.Printf("type of response is %T\n", response)

	//caller's responsibility to close the connection
	defer response.Body.Close()

	// read the response
	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)

	}
	contents := string(databytes)

	fmt.Println(contents)

}
