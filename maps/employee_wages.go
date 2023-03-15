package main

import "fmt"

func main() {

	employeeSalary1 := make(map[string]int)

	employeeSalary1["John Wick"] = 115000
	employeeSalary1["Jammy Fox"] = 135000
	employeeSalary1["Chris Rock"] = 125000
	fmt.Println("Group1 Wages\n ", employeeSalary1)

	employeeSalary := map[string]int{"Jet Lee": 65000}
	employeeSalary["Jackie Chan"] = 34000
	fmt.Println("Group2 Wages\n ", employeeSalary)

}
