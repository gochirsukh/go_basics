package main

import "fmt"

func main() {
	byteArr := [5]byte{104, 101, 108, 108, 111}
	str := string(byteArr[:])
	fmt.Println(str)
}
