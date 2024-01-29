package main

import "fmt"

func main9() {
	var Nbr11 int
	var Nbr21 int
	var Nbr31 int
	var Nbr12 int
	var Nbr22 int
	var Nbr32 int

	fmt.Print("Please enter the number times each variable occur and the value of the two equations (x, y, =z), ")

	fmt.Print("first number equation 1:")
	fmt.Scanln(&Nbr11)
	fmt.Print("Second number equation 1:")
	fmt.Scanln(&Nbr21)
	fmt.Print("value equation 1:")
	fmt.Scanln(&Nbr31)
	fmt.Print("first number equation 2:")
	fmt.Scanln(&Nbr12)
	fmt.Print("second number equation 2:")
	fmt.Scanln(&Nbr22)
	fmt.Print("value equation 2:")
	fmt.Scanln(&Nbr32)

	e1 := [3]int{Nbr11, Nbr21, Nbr31} // {a1,b1,c1}
	e2 := [3]int{Nbr12, Nbr22, Nbr32} // {a2,b2,c2} a2= e2[0]

	var e1a2 [3]int
	var e2a1 [3]int

	for i := 0; i < 3; i++ {
		e1a2[i] = e1[i] * e2[0]
	}
	fmt.Printf("%v\n", e1a2)
	for i := 0; i < 3; i++ {
		e2a1[i] = e2[i] * e1[0]
	}
	fmt.Printf("%v\n", e2a1)
	tmp1 := e1a2[1] - e2a1[1]
	tmp2 := e1a2[2] - e2a1[2]
	fmt.Print("Y equals:")
	fmt.Println(tmp2 / tmp1)
	Value1 := tmp2 / tmp1
	fmt.Print("X equals:")
	fmt.Println(Nbr31 - Value1*Nbr21)
}
