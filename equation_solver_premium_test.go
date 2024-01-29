package main

import (
	"fmt"
	"testing"
)

func TestE5(t *testing.T) {
	var parameters = [][]int{
		{1, 1, 1, 1, 1, 15},
		{2, 3, 3, 3, 3, 44},
		{0, 2, 2, 2, 2, 28},
		{2, 0, 3, 4, 5, 52},
		{3, 4, 2, 2, 6, 55},
	}
	printAllEquations(parameters)
	answers := solveE(parameters)
	correct := verifyAnswers(answers, parameters)
	fmt.Println(correct, answers)
}

//func TestE2Failed1(t *testing.T) {
	var parameters = [][]int{
		{-62, 186, 21452},
		{-23, 69, 7958},
	}
	printAllEquations(parameters)
	answers := solveE(parameters)
	correct := verifyAnswers(answers, parameters)
	fmt.Println(correct, answers)
	/*
		       -62X₀      +186X₁ =      21452
		       -23X₀       +69X₁ =       7958
		false [-493393 1]
	*/
}
func TestE3Failed1(t *testing.T) {
	var parameters = [][]int{
		{196, 6, 87, 34069},
		{0, 33, -6, 2751},
		{62, 102, -33, 20603},
	}
	printAllEquations(parameters)
	answers := solveE(parameters)
	correct := verifyAnswers(answers, parameters)
	fmt.Println(correct, answers)
	/*
			       +196X₀        +6X₁       +87X₂ =      34069
		        +0X₀       +33X₁        -6X₂ =       2751
		       +62X₀      +102X₁       -33X₂ =      20603
		false [81.91651233387418 119.5105543499704 198.80804892483724]
	*/
}

func TestE3(t *testing.T) {
	var parameters = [][]int{
		{2, 8, 9, 45},
		{3, 7, -1, 14},
		{-1, -2, 1, -2},
	}
	printAllEquations(parameters)
	answers := solveE(parameters)
	correct := verifyAnswers(answers, parameters)
	fmt.Println(correct, answers)
}

func TestE6(t *testing.T) {
	var parameters = generateRandomParametersAsInput(6, 1, 10)
	printAllEquations(parameters)
	answers := solveE(parameters)
	correct := verifyAnswers(answers, parameters)
	fmt.Println(correct, answers)
}

func TestEMany2(t *testing.T) {
	var totalTrue, totalFalse int
	totalTest := 1000000
	for i := 0; i < totalTest; i++ {
		var parameters = generateRandomParametersAsInput(2, -100, 200)
		//printAllEquations(parameters)
		answers := solveE(parameters)
		correct := verifyAnswers(answers, parameters)
		if correct {
			totalTrue++
		} else {
			totalFalse++
			printAllEquations(parameters)
			fmt.Println(correct, answers)
		}
	}
	fmt.Println(totalTest, totalTrue, totalFalse)
}

func TestEMany3(t *testing.T) {
	var totalTrue, totalFalse int
	resultCh := make(chan int, 10000000)
	totalTest := 100000000
	for i := 0; i < totalTest; i++ {
		go func() {
			var parameters = generateRandomParametersAsInput(3, -100, 200)
			//printAllEquations(parameters)
			answers := solveE(parameters)
			correct := verifyAnswers(answers, parameters)
			if correct {
				resultCh <- 1
			} else {
				resultCh <- 1
				printAllEquations(parameters)
				fmt.Println(correct, answers)
			}
		}()
	}

	for {
		result := <-resultCh
		if result == 1 {
			totalTrue++
		} else {
			totalFalse++
		}
		if totalFalse+totalTrue == 10000000 {
			break
		}
	}
	fmt.Println(totalTest, totalTrue, totalFalse)
}

func TestEMany4(t *testing.T) {
	var totalTrue, totalFalse int
	totalTest := 10000
	for i := 0; i < totalTest; i++ {
		var parameters = generateRandomParametersAsInput(4, -100, 200)
		//printAllEquations(parameters)
		answers := solveE(parameters)
		correct := verifyAnswers(answers, parameters)
		if correct {
			totalTrue++
		} else {
			totalFalse++
			printAllEquations(parameters)
			fmt.Println(correct, answers)
		}
	}
	fmt.Println(totalTest, totalTrue, totalFalse)
}

func TestEMany3NoGoroutines(t *testing.T) {
	var totalTrue, totalFalse int
	totalTest := 10000000
	for i := 0; i < totalTest; i++ {
		var parameters = generateRandomParametersAsInput(3, -100, 200)
		//printAllEquations(parameters)
		answers := solveE(parameters)
		correct := verifyAnswers(answers, parameters)
		if correct {
			totalTrue++
		} else {
			totalFalse++
			printAllEquations(parameters)
			fmt.Println(correct, answers)
		}
	}
	fmt.Println(totalTest, totalTrue, totalFalse)
}
