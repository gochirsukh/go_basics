package main

import (
	"fmt"
	"time"
)

func displayDate(message string, format string) {
	fmt.Println(message, time.Now().Format(format))
}

func main() {
	displayDate("Current Date and time", "19:06:03, 2023-March-26 Sun")

}
