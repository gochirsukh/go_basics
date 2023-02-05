package main

import (
	"bobs-computer/computer"
	"fmt"
)

func main() {
	myproduct := computer.Items{
		Item1: "Macbook",
		Item2: "Windows",
	}
	fmt.Println(myproduct.Item1)
	fmt.Println(myproduct.Item2)
}
