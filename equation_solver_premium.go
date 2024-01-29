package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//func main() {

	fmt.Println("Please enter numbers of unknown integers:")
	var NumberOfUnknown int
	_, _ = fmt.Scanln(&NumberOfUnknown)

	parameters := getAllEParameters(NumberOfUnknown)

	var finalAnswer []float64
	finalAnswer = solveE(parameters)
	fmt.Println("Here's your answer based on the order you entered the integers at the beginning, if your equation was impossible from the beginning, then the output will not be correct either:")
	fmt.Println(finalAnswer)
}

func Solve2(parameters [][]int) []float64 {
	var e1a2 [3]float64
	var e2a1 [3]float64
	for i := 0; i < 3; i++ {
		e1a2[i] = float64(parameters[0][0] * parameters[1][i])
		checkOverflow(e1a2[i], float64(parameters[1][i]), float64(parameters[0][0]))
		e2a1[i] = float64(parameters[1][0] * parameters[0][i])
		checkOverflow(e2a1[i], float64(parameters[0][i]), float64(parameters[1][0]))
	}
	tmp1 := e1a2[1] - e2a1[1]
	tmp2 := e1a2[2] - e2a1[2]
	var x, y float64
	if tmp1 == 0 && tmp2 == 0 {
		y = 1
		x = (e1a2[2] - e1a2[1]) / e1a2[0]
		return []float64{x, y}
	} else if tmp1 == 0 && tmp2 != 0 {
		fmt.Println("Equation has no solution")
		os.Exit(0)
	}
	y = tmp2 / tmp1
	x = (float64(parameters[0][2]) - y*float64(parameters[0][1])) / float64(parameters[0][0])
	return []float64{x, y}
}

func GetOneEquationParameters() []int {
	var e1parameters string
	fmt.Println("Enter occurrence value of equations:")
	_, err := fmt.Scanln(&e1parameters)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	tmpParameters := strings.Split(e1parameters, ",")

	var intParas []int
	for i := 0; i < len(tmpParameters); i++ {
		tmpInt, err := strconv.Atoi(tmpParameters[i])
		if err != nil {
			fmt.Println("wrong input format", err.Error())
			return nil
		}
		intParas = append(intParas, tmpInt)
	}
	return intParas
}

func getAllEParameters(numberOfUnknown int) [][]int {
	var parameters [][]int
	for i := 0; i < numberOfUnknown; i++ {
		parameters = append(parameters, GetOneEquationParameters())
	}
	return parameters
}

func multiplier(parameters [][]int) [][]int {
	var originalAx []int
	for k := 0; k < len(parameters); k++ {
		originalAx = append(originalAx, parameters[k][0])
	}

	for k := 0; k < len(parameters); k++ {
		coefficient1 := originalAx[k]
		if coefficient1 == 0 {
			continue
		}
		for i := 0; i < len(parameters); i++ {
			if i == k {
				continue
			}
			for j := 0; j < len(parameters[i]); j++ {
				tmp1 := parameters[i][j] * coefficient1
				checkOverflow(float64(tmp1), float64(parameters[i][j]), float64(coefficient1))
				parameters[i][j] = tmp1
			}
		}
	}
	return parameters
}

func integerSliceSubtract(e1, e2 []int) []int {
	var parameters []int
	for i := 1; i < len(e1); i++ {
		subtractedResult := e1[i] - e2[i]
		parameters = append(parameters, subtractedResult)
	}
	return parameters
}

func reduceExcessInt(parameters [][]int) [][]int {
	var reducedInt [][]int
	for i := 0; i < len(parameters)-1; i++ {
		reducedInt2 := integerSliceSubtract(parameters[i], parameters[i+1])
		reducedInt = append(reducedInt, reducedInt2)
	}
	return reducedInt
}

func solveE(inputParameters [][]int) []float64 {
	inputParameters = smallestEfficientByGCD(inputParameters)

	if len(inputParameters) > 2 {
		parameters := makeCopyParameter(inputParameters)
		var coEfficiency0Equation [][]int
		var lParas [][]int
		for i := 0; i < len(parameters); i++ {
			if parameters[i][0] == 0 {
				coEfficiency0Equation = append(coEfficiency0Equation, parameters[i][1:])
			} else {
				lParas = append(lParas, parameters[i])
			}
		}
		lParas = multiplier(lParas)
		reducer := reduceExcessInt(lParas)
		reducer = smallestEfficientByGCD(reducer)
		//fmt.Println("after reduceExcessInt", reducer)
		if len(coEfficiency0Equation) > 0 {
			reducer = append(reducer, coEfficiency0Equation...)
		}
		prevFunctions := solveE(reducer) // [2,3,4,5]
		x := calculateVar(inputParameters, prevFunctions)
		newAnswers := []float64{x}
		newAnswers = append(newAnswers, prevFunctions...)
		return newAnswers
	}

	a1 := inputParameters[0][0]
	b1 := inputParameters[0][1]
	a2 := inputParameters[1][0]
	b2 := inputParameters[1][1]
	if a1 == 0 || b1 == 0 || a2 == 0 || b2 == 0 {
		return solve2SpecialSit(inputParameters)
	}
	return Solve2(inputParameters)
}

func solve2SpecialSit(parameters [][]int) []float64 {
	var x, y float64
	a1 := float64(parameters[0][0])
	b1 := float64(parameters[0][1])
	c1 := float64(parameters[0][2])
	a2 := float64(parameters[1][0])
	b2 := float64(parameters[1][1])
	c2 := float64(parameters[1][2])

	if a1 == 0 && a2 == 0 && b1 == 0 && b2 == 0 {
		return []float64{1, 1}
	}

	if a1 == 0 && a2 == 0 {
		if b1 != 0 {
			y = c1 / b1
			return []float64{1, y}
		}
		y = c2 / b2
		return []float64{1, y}
	}
	if b1 == 0 && b2 == 0 {
		if a1 != 0 {
			x = c1 / a1
			return []float64{x, 1}
		}
		x = c2 / a2
		return []float64{x, 1}
	}

	if a1 == 0 && b1 == 0 {
		x = 1
		y = (c2 - a2) / b2
		return []float64{x, y}
	}

	if a2 == 0 && b2 == 0 {
		y = 1
		x = (c1 - b1) / a1
		return []float64{x, y}
	}

	if a1 == 0 {
		y = c1 / b1
		x = (c2 - b2*y) / a2
		return []float64{x, y}
	}
	if a2 == 0 {
		y = c2 / b2
		x = (c1 - b1*y) / a1
		return []float64{x, y}
	}
	if b1 == 0 {
		x = c1 / a1
		y = (c2 - a2*x) / b2
		return []float64{x, y}
	} else { //b2==0
		x = c2 / a2
		y = (c1 - a1*x) / b1
		return []float64{x, y}
	}
}

func makeCopyParameter(parameters [][]int) [][]int {
	var newV [][]int
	for i := 0; i < len(parameters); i++ {
		var tmpSlice1 []int
		for j := 0; j < len(parameters[i]); j++ {
			tmpSlice1 = append(tmpSlice1, parameters[i][j])
		}
		newV = append(newV, tmpSlice1)
	}
	return newV
}

func calculateVar(parameters [][]int, answers []float64) float64 {
	for i := 0; i < len(parameters); i++ {
		if parameters[i][0] == 0 {
			continue
		}
		var sum float64
		for j := 1; j < len(parameters[i])-1; j++ {
			sum = sum + float64(parameters[i][j])*answers[j-1]
		}
		x := (float64(parameters[i][len(parameters[i])-1]) - sum) / float64(parameters[i][0])
		return x
	}
	return 1
}

func smallestEfficientByGCD(reducer [][]int) [][]int {
	for i := 0; i < len(reducer); i++ {
		g := GetGcd(reducer[i])
		if g > 1 {
			for j := 0; j < len(reducer[i]); j++ {
				reducer[i][j] = reducer[i][j] / g
			}
		}
	}
	return reducer
}

func generateRandomParametersAsInput(numberOfUnknown, min, max int) [][]int {
	var generatedRandomAnswers []int
	for i := 0; i < numberOfUnknown; i++ {
		generatedRandomAnswers = append(generatedRandomAnswers, generateOneRandomInt(min, max))
	}
	var allEquationParameters [][]int
	for i := 0; i < numberOfUnknown; i++ {
		var oneEquationParameters []int
		var sum int
		for j := 0; j < numberOfUnknown; j++ {
			coefficient := generateOneRandomInt(min, max)
			oneEquationParameters = append(oneEquationParameters, coefficient)
			sum = sum + coefficient*generatedRandomAnswers[j]
		}
		oneEquationParameters = append(oneEquationParameters, sum)
		allEquationParameters = append(allEquationParameters, oneEquationParameters)
	}
	return allEquationParameters
}

func generateOneRandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func verifyAnswers(answers []float64, allEquationParameters [][]int) bool {
	numberOfEquation := len(allEquationParameters)
	for i := 0; i < numberOfEquation; i++ {
		var calculatedSumFromAnswer float64
		for j := 0; j < len(allEquationParameters[i])-1; j++ {
			calculatedSumFromAnswer = calculatedSumFromAnswer + float64(allEquationParameters[i][j])*answers[j]
		}
		sumFromInput := float64(allEquationParameters[i][len(allEquationParameters[i])-1])
		if math.Abs(sumFromInput-calculatedSumFromAnswer) > 0.01 {
			fmt.Println(math.Abs(sumFromInput-calculatedSumFromAnswer), sumFromInput, calculatedSumFromAnswer)
			return false
		}
	}
	return true
}

func printAllEquations(parameters [][]int) {
	for i := 0; i < len(parameters); i++ {
		for j := 0; j < len(parameters[i])-1; j++ {
			coefficient := parameters[i][j]
			if coefficient >= 0 {
				fmt.Printf("%10sX%c", "+"+strconv.Itoa(coefficient), 0x2080+j)
			} else {
				fmt.Printf("%10sX%c", strconv.Itoa(coefficient), 0x2080+j)
			}
		}
		fmt.Printf(" = %10v\n", parameters[i][len(parameters[i])-1])
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func GetGcd(n []int) int {
	g := n[0]
	for i := 1; i < len(n)-1; i++ {
		g = gcd(g, n[i])
	}
	return g
}

func checkOverflow(result, input1, input2 float64) {
	if result == 0 {
		return
	}
	if math.Abs(result/input1-input2) > 0.01 {
		fmt.Printf("overflow found: result=%v input1=%v input2=%v\n", result, input1, input2)
		os.Exit(0)
	}
}
