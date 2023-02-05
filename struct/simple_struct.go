package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var p Person
	fmt.Println(p)

	p1 := Person{"John", "Wick", 43}
	fmt.Println("Person1", p1)

	p2 := Person{
		FirstName: "Arnold",
		LastName:  "Schwartsnegger",
		Age:       45,
	}

	fmt.Println("Person2: ", p2)

	p3 := Person{FirstName: "Jackie"}
	fmt.Println("Person3: ", p3)
}
