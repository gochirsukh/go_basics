package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		// panic(e)
		fmt.Println("File not found:\n", e)
	}

}

func main() {
	file, err := os.ReadFile("README.md")
	check(err)
	fmt.Println(string(file))
}
