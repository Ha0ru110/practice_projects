package main

import "fmt"

func main2() {
	var guess int
	number := 35
	for guess != number {
		fmt.Println("Please enter your guess")
		fmt.Scanln(&guess)

		if guess > number {
			fmt.Println("too high")
		}
		if number > guess {
			fmt.Println("too low")
		}
		if guess == number {
			fmt.Println("well done, game over")
		}

	}
}
