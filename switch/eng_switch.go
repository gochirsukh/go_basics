package main

import "fmt"

func main() {

	// Take input
	fmt.Println("Enter name of a team member")
	var name string
	fmt.Scan(&name)

	switch name := name; name {
	case "Gansukh":
		fmt.Print("He is SRE")
	case "McGovern":
		fmt.Print("He is an architect")
	case "Angie":
		fmt.Print("She is a manager")
	default:
		fmt.Printf("We don't know %v\n", name)
	}

}
