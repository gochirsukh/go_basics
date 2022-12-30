package main

// https://gobyexample.com/switch

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	}
	myos := runtime.GOOS
	fmt.Println("OS is: ", myos)
}
