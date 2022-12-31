package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter name")
	var name string

	fmt.Scanln(&name)

	switch name := name; name {
	case "Gansukh":
		fmt.Println("He is a SRE")
	case "Chris":
		fmt.Println("He is an analyst")
	default:
		fmt.Println("Who knows?")
	}

}
