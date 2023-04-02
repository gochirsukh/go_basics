package main

import (
	"fmt"

	"oakwood.com/worker-app/person_attr"
)

func main() {

	p1 := person_attr.Demography{
		Fname: "John",
		Lname: "Wick",
		Age:   45,
	}
	fmt.Printf(" Firstname: %v\n Lastname: %v\n Age: %v\n", p1.Fname, p1.Lname, p1.Age)

}
