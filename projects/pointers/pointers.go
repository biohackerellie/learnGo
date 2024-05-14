package main

import "fmt"

func main() {
	age := 31
	var agePointer *int

	agePointer = &age
	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Adult years: ", age)
}

func getAdultYears(age *int) int {
	*age = *age - 18
	return *age
}
