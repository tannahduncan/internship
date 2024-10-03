// json to Go
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
	// JSON data as a string
	jsonData := `{"name":"Alice","age":30}`

	// Variable to hold the unmarshalled data
	var p Person

	// Unmarshalling JSON into Go struct
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Printing Go struct fields
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}
