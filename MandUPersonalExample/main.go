package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Creating an instance of Person
	p := Person{
		Name: "Alice",
		Age:  30,
	}

	// Marshaling Go struct into JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Printing JSON string
	fmt.Println(string(jsonData))
}
