// Resources:
// https://go.dev/blog/maps
// https://gobyexample.com/maps

package main

import "fmt"

func main() {

	// (map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)
	// ^^ This is a bealean that checks whether K2 exists or not

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
