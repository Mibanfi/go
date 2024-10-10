package main

import "fmt"

func main() {
	var input int
	var isSameOutput bool
	input = 0
	isSameOutput = true
	for isSameOutput {
		isSameOutput = (function1(input) == function2(input))
		input++
	}
	input--
	fmt.Print("Error on input ", input, ": function1 returns ", function1(input), " function 2 returns ", function2(input), "\nAborting.\n")
}

func function1(n int) int {
	return n
}

func function2(n int) int {
	return n
}
