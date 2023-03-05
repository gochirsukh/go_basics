package main

import (
	"fmt"

	"example.com/human_main_dir/person_module_dir"
)

func main() {

	p1 := person_module_dir.PerAttr{

		Fname: "John",
		Lname: "Wick",
	}

	fmt.Println(p1.Fname, p1.Lname)
}
