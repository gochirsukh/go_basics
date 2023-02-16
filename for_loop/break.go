// resource
// https://www.digitalocean.com/community/tutorials/using-break-and-continue-statements-when-working-with-loops-in-go
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("Loop number: ", i)
		if i == 4 {
			fmt.Println("There is a break statement at iteration 4, breaking out of the loop")
			break
		}
	}
	fmt.Println("Exiting the program")
}
