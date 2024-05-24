package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)
	//adding a key and value
	//languages["key"] = "value"
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])
   
	//removing a key
	delete(languages, "RB")

	fmt.Println("List of all languages: ", languages)

	//loop through map

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value) //%v is used to print the value
	}

}
