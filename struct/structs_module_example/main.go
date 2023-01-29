// First run "go mod init structs_module_example"
// Example from https://golangbot.com/structs/

package main

import (
	"fmt"

	"structs_module_example/computer"
)

func main() {
	specs := computer.Spec{
		Maker: "Apple",
		Model: "MacBook Pro",
		Price: 2000,
	}
	fmt.Println("Maker:", specs.Maker)
	fmt.Println("Model:", specs.Model)
	fmt.Println("Price: $", specs.Price)

}
