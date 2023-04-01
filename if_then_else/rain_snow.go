package main

// Application demonstrates or || operator

import "fmt"

func raining() bool {
	fmt.Println("Check if it is raining now")
	return true
}

func snowing() bool {
	fmt.Println("Check if it is snowing now")
	return true
}

// if first condition is true, then the second function doesnt' need to be called
func main() {
	if raining() || snowing() {
		fmt.Println("Stay indoors")
	} else {
		fmt.Println("Go outdoors")
	}

}
