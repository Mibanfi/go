package main

import (
	"biscuit"
	"fmt"
)

func main() {
	var numbers []string
	numbers = biscuit.Input(3, 3, []string{"number1", "number2", "number3"})
	fmt.Println(numbers)
}
