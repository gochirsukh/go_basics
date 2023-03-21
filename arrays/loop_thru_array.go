package main

import "fmt"

func main() {

	arrOfString := [4]string{"orange", "apple", "peach", "melon"}

	for i := 0; i < len(arrOfString); i++ {
		fmt.Println(arrOfString[i])

	}

}
