package main

import "fmt"

func main() {

	fmt.Println("Enter name")
	var fname string
	fmt.Scan(&fname)

	switch vname := fname; vname {
	case "Gansukh":
		fmt.Println("He is an architect and SRE")
	case "Williams":
		fmt.Println("He is an analyst")
	default:
		fmt.Printf("We don't know who %v is\n", fname)
	}

}
