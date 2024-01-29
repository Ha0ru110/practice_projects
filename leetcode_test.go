package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {

	var num1int int
	var num2int int
	var n int
	var saveData []int
	var answer int
	var answer2 int

	if len(num1) != len(num2) {
		num1, num2 = addZeroToShortString(num1, num2)
	}
	n = len(num1)
	for i := n; i >= 0; i-- {
		num1int, _ = strconv.Atoi(string(num1[i]))
		num2int, _ = strconv.Atoi(string(num2[i]))
		answer, answer2 = addition(num1int, num2int, answer2)
		saveData = append(saveData, answer)
	}
	if answer2 != 0 {
		saveData = append(saveData, answer2)
	}

	return makeFinalAnswer(saveData)
}

func addition(num1int, num2int, prerest int) (int, int) {
	sum := num1int + num2int + prerest
	return sum % 10, sum / 10
}
func addZeroToShortString(num1, num2 string) (string, string) {
	var result string

	if len(num1) > len(num2) {
		result = addHeadZero(len(num1)-len(num2), num2)
		return result, num1
	}
	if len(num1) < len(num2) {
		result = addHeadZero(len(num2)-len(num1), num1)
		return result, num2
	}
	return num1, num2
}

func addHeadZero(times int, tailString string) string {
	var sliceString string
	for i := 0; i < times; i++ {
		sliceString += "0"
	}
	return sliceString + tailString
}
func makeFinalAnswer(saveData []int) string {
	var finalAnswer string
	for i := len(saveData); i >= 0; i-- {
		currentString := strconv.Itoa(saveData[i])
		finalAnswer += currentString
	}
	return finalAnswer
}
