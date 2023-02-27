package main

import "fmt"

func main() {

	var v1 = 400
	var v2 = 700

	if v1 == 400 {
		fmt.Printf("Value of v1 is %v\n", v1)
		if v2 == 700 {
			fmt.Printf("Value of v2 is %v\n", v2)
		}
	}
}
