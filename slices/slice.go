// Resource:
// https://gobyexample.com/slices
package main

import "fmt"

func main() {
	s := make([]string, 4)
	fmt.Println("Empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	fmt.Println("List them all:", s)
	fmt.Println("Get:", s[2])
	fmt.Println("Len:", len(s))

}
