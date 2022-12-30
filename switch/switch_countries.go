package main

import (
	"fmt"
)

func main() {

	country := "Mongolia"

	switch country := "Mongolia"; country {
	case "Mongolia":
		fmt.Println("horses")

	case "China":
		fmt.Println("rice")
	}
}
