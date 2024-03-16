package main

import (
	"fmt"
)

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"canoe", "piano", "skiing"}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	hobbies2 := hobbies[1:]
	fmt.Println(hobbies2)
	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	goals := []string{"Get better at programming", "start a business"}
	fmt.Println(goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	goals[1] = "understand interfaces"
	goals = append(goals, "build an API")
	fmt.Println(goals)
	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{{"first product", "A first product", 5.99}, {"second product", "a second product", 6.99}}
	fmt.Println(products)
	newProduct := Product{"third product", "a third product", 12.99}
	products = append(products, newProduct)
	fmt.Println(products)
}
