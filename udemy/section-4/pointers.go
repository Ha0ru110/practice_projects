package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)
	agePointer := &age
	fmt.Println("Age: ", *agePointer)
	getAdultYears(agePointer)
	fmt.Println(age)

}

func getAdultYears(agePointer *int) int {
	*agePointer -= 18
	return *agePointer
}
