package main

import . "fmt"

func main() {
	var age int
	_, _ = Scan(&age)
	agePointer := &age
	Println("Age: ", *agePointer)
	getAdultYears(agePointer)
	Println(age)

}

func getAdultYears(agePointer *int) int {
	*agePointer -= 18
	return *agePointer
}
