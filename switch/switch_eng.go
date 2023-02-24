package main

import "fmt"

func main() {

	//Print out the question
	fmt.Println("Enter name")
	//Declare var
	var fname string
	//Take name
	fmt.Scan(&fname)

	//Switch
	switch vname := fname; vname {
	case "Gansukh":
		fmt.Println("He is an architect and SRE")
	case "Williams":
		fmt.Println("He is an analyst")
	default:
		fmt.Printf("We don't know who %v is\n", fname)
	}

}
