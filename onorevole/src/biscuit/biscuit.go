// A package with all the functions I use over and over for input management, error handling, debugging etc.
// The name "biscuit" doesn't mean anything.
// Well, actually, it's because this package contains "cookie-cutter" features, and I wanted to avoid confusion with internet cookies. So, biscuit it is.

package biscuit

import (
	"fmt"
	"os"
	"strconv"
)

func Input(minArgs int, maxArgs int, descriptors []string) []string {
	// Setup variables
	var arguments []string // Used to store the arguments which will be returned
	if maxArgs == 0 {
		arguments = make([]string, len(os.Args))
	} else {
		arguments = make([]string, maxArgs)
	}
	var index int // Used in for cycles

	// Ensure we have the at least the minimum number of arguments. Setting minArgs to 0 ignores the check.
	if minArgs != 0 && len(os.Args) < minArgs {
		fmt.Println("Not enough arguments!")
		fmt.Println("We need at least", minArgs, "arguments, in the following format:", descriptors)
		fmt.Println("Arguments must be typed after calling the command, separated by spaces.")
		fmt.Println("Let's insert the missing arguments.")
		for index = len(os.Args); index < minArgs; index++ {
			fmt.Print(descriptors[min(index, len(descriptors)-1)], ": ")
			fmt.Scan(&arguments[index])
		}
	}

	// Ensure we have the no more than the maximum number of arguments. Setting maxArgs to 0 ignores the check.
	if maxArgs != 0 && len(os.Args) > maxArgs {
		fmt.Println("Too many arguments!")
		fmt.Println("We need no more than", maxArgs, "arguments, in the following format:", descriptors)
		fmt.Println("Arguments must be typed after calling the command, separated by spaces.")
		fmt.Println("The arguments in excess will be ignored.")
	}

	// Copy existing arguments from command line
	for index = 0; index < min(len(os.Args), maxArgs); index++ {
		arguments[index] = os.Args[index+1]
	}

	// Return results
	return arguments
}

func InputInt(minArgs int, maxArgs int, descriptors []string) []int {
	// Setup variables
	var argString []string // Stores arguments as string type
	var argInt []int       // Stores arguments as int type
	var err error          // Used to store error messages
	var index int          // Used in for cycles
	argString = Input(minArgs, maxArgs, descriptors)
	argInt = make([]int, len(argString))

	// Convert all arguments
	for index = 0; index < len(argString); index++ {
		argInt[index], err = strconv.Atoi(argString[index])
		for err != nil {
			fmt.Println("Error:", argString[index], "is not a number. Please input a new value.")
			fmt.Print(descriptors[index], ": ")
			fmt.Scan(&argString[index])
			argInt[index], err = strconv.Atoi(argString[index])
		}
	}

	// Return all arguments
	return argInt
}
