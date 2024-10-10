package main

import (
	"fmt"
	"plutopos"
)

func main() {
	var age, i uint64
	fmt.Print("Please input your updated age: ")
	fmt.Scan(&age)
	for i = 1; i < age; i++ {
		fmt.Print(plutopos.NumToWord(i), ", ")
	}
	fmt.Println("\n..aaaaaaand...")
	fmt.Print(plutopos.NumToWord(age), "! HAPPY BIRTHDAY!!")
	fmt.Println("\n\nI love you :D")
}
