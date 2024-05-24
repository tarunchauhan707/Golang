package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://samrat.dev:8080/learn?course=golang&paymentid=123"

func main() {
	fmt.Println("Welcome to url package in golang")

	fmt.Println(myurl)

	//parsing url
    result, _:=url.Parse(myurl)
    
	//information about url
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//query
	qparams:= result.Query()
	fmt.Printf("The type of query params are :%T\n", qparams)
	fmt.Println(qparams["course"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Parameter is", val)
	}

	partsOfurl:= &url.URL{
		Scheme: "https",
		Host : "samrat.dev",
		Path : "/learn",
		RawPath: "user=samrat",
	}
	anotherURL := partsOfurl.String()
	fmt.Println(anotherURL)
}
