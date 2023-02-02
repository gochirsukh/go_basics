package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Outer loop: %d\t The inner loop: %d\n ", i, j)
		}
	}
}
