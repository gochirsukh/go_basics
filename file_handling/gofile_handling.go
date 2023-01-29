package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	body, err := os.ReadFile("file.txt")
	if err != nil {
		log.Fatalf("Unable to read file! Does it exist? ...", err)
	}
	fmt.Println(string(body))
}
