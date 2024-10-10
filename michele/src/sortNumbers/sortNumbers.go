package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Debugging variables
	var err error
	var argsImport []string
	argsImport = os.Args

	// Initializing
	var inputNumbers, sortedNumbers []int

	// Collecting input
	if len(argsImport) > 2 {
		inputNumbers = make([]int, len(argsImport)-1)
		for i := 0; i < len(inputNumbers); i++ {
			inputNumbers[i], err = strconv.Atoi(argsImport[i+1])
			for err != nil {
				fmt.Println("Error:", argsImport[i+1], "is not a number!")
				fmt.Println("Please input a number to replace it, or Ctrl+C to abort.")
				fmt.Scan(&argsImport[i+1])
				inputNumbers[i], err = strconv.Atoi(argsImport[i+1])
			}
		}
	} else {
		inputNumbers = make([]int, 3)
		fmt.Print("Numero 1: ")
		fmt.Scan(&inputNumbers[0])
		fmt.Print("Numero 2: ")
		fmt.Scan(&inputNumbers[1])
		fmt.Print("Numero 3: ")
		fmt.Scan(&inputNumbers[2])
	}

	// Returning output
	sortedNumbers = sortNumbers(inputNumbers)
	fmt.Println(sortedNumbers)
}

func sortNumbers(numbers []int) []int {
	// Initializing
	var index1, index2 int

	// Operating
	for index1 = 0; index1 < len(numbers); index1++ {
		for index2 = index1; index2 < len(numbers); index2++ {
			if numbers[index2] < numbers[index1] {
				numbers[index2], numbers[index1] = numbers[index1], numbers[index2]
			}
		}
	}

	// Returning
	return numbers
}
