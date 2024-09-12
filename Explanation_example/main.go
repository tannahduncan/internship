package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type PersonV1 struct {
	FirstName string
	LastName  string
	Age       int
}

type PersonV2 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age,omitempty"`
}

type Job struct {
	Title      string `json:"title"`
	PayGrade   int    `json:"paygrade"`
	Supervisor string `json:"supervisor"`
}

type PersonV3 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Job       Job    `json:"job,omitempty"`
}

type PersonV4 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Job       *Job   `json:"job,omitempty"`
}

type PersonV5 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Jobs      []Job  `json:"jobs,omitempty"`
}

func main() {

	myPerson1 := PersonV1{FirstName: "Brian", LastName: "DeBusk", Age: 56}

	fmt.Printf("Person1: %+v\n", myPerson1)

	//json.MarshalIndent: Converts an object to JSON with indentation
	//Error Handling: Logs and stops execution if marshaling fails

	jsonBytes, err := json.MarshalIndent(myPerson1, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV1: \n%s\n", string(jsonBytes))
	// ----------------------------------

	myPerson2 := PersonV2{FirstName: "Brian", LastName: "DeBusk", Age: 56}

	fmt.Printf("Person2: %+v\n", myPerson2)

	jsonBytes, err = json.MarshalIndent(myPerson2, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV2: \n%s\n", string(jsonBytes))
	// ----------------------------------

	myPerson2a := PersonV2{FirstName: "Brian", LastName: "DeBusk"}

	fmt.Printf("Person2a: %+v\n", myPerson2a)

	jsonBytes, err = json.MarshalIndent(myPerson2a, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV2a: \n%s\n", string(jsonBytes))

	// ----------------------------------

	myJob := Job{Title: "Developer", PayGrade: 5, Supervisor: "Joe"}
	myPerson3 := PersonV3{FirstName: "Brian", LastName: "DeBusk", Age: 56, Job: myJob}

	//myJob is nested into myPerson3

	fmt.Printf("Person3: %+v\n", myPerson3)

	jsonBytes, err = json.MarshalIndent(myPerson3, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV3: \n%s\n", string(jsonBytes))
	// ----------------------------------

	// Subtle nuance here... it will still create an empty "Job" data element inside this structure
	myPerson3a := PersonV3{FirstName: "Brian", LastName: "DeBusk", Age: 56}

	fmt.Printf("Person3a: %+v\n", myPerson3a)

	jsonBytes, err = json.MarshalIndent(myPerson3a, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV3a: \n%s\n", string(jsonBytes))
	// ----------------------------------

	myPerson4 := PersonV4{FirstName: "Brian", LastName: "DeBusk", Age: 56}

	fmt.Printf("Person4: %+v\n", myPerson4)

	jsonBytes, err = json.MarshalIndent(myPerson4, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV4: \n%s\n", string(jsonBytes))
	// ----------------------------------

	myPerson4a := PersonV4{FirstName: "Brian", LastName: "DeBusk", Age: 56}
	// Must have this... or it creates a nil pointer reference
	myPerson4a.Job = new(Job)
	myPerson4a.Job.Title = "Engineer"
	myPerson4a.Job.PayGrade = 6
	myPerson4a.Job.Supervisor = "Joe"

	fmt.Printf("Person4a: %+v\n", myPerson4a)

	jsonBytes, err = json.MarshalIndent(myPerson4a, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV4a: \n%s\n", string(jsonBytes))
	// ----------------------------------

	myPerson5 := PersonV5{FirstName: "Brian", LastName: "DeBusk", Age: 56}

	fmt.Printf("Person5: %+v\n", myPerson5)

	jsonBytes, err = json.MarshalIndent(myPerson5, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV5: \n%s\n", string(jsonBytes))
	// ----------------------------------

	myPerson5.Jobs = append(myPerson5.Jobs, Job{Title: "Developer", PayGrade: 5, Supervisor: "Joe"})
	myPerson5.Jobs = append(myPerson5.Jobs, Job{Title: "Engineer", PayGrade: 6, Supervisor: "Joe"})

	jsonBytes, err = json.MarshalIndent(myPerson5, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Marshaled PersonV5 with jobs added: \n%s\n", string(jsonBytes))
	// ----------------------------------

	// yaml.Marshal: Converts data to YAML format, yaml: more compact, easier to read

	yamlBytes, err := yaml.Marshal(myPerson5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("YAML Marshaled PersonV5 with jobs added: \n%s\n", string(yamlBytes))

	// json.Unmarshal: Converts JSON data to a Go map

	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Generic Unmarshaling:\n%+v\n", result)
	// In a dynamically cast world, this would work: result["jobs"][1]["title"]
	// and even use a dot notation: "jobs.1.title"
	// But in Golang, we have to statically cast at each level
	fmt.Printf("Reference: %v\n", result["jobs"].([]interface{})[1].(map[string]interface{})["title"])
}
