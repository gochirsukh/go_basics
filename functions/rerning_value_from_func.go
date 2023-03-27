package main

import "fmt"

func addNum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	result := addNum(3, 4)
	fmt.Println(result)
}
