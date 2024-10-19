package main

import "fmt"

func main() {
	age := 28 //regular variable

	agePointer := &age

	fmt.Println("Age: ", *agePointer)
	editAgeToAdultYears(agePointer)
	fmt.Println("Age since being adult(18)", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
