package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i == 3 {
			fmt.Println("Loop continues")
			continue
		}
		fmt.Printf("GoLang is easy %v\n", i)
	}
	fmt.Println("Exiting the loop")
}
