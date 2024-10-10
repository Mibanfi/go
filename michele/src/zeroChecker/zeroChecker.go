package main

import "fmt"

func main() {
	var number, zeroes int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&number)
	fmt.Print("\nInserisci il numero di zeri finali che dovrebbe avere: ")
	fmt.Scan(&zeroes)
	fmt.Print("\n")
	if checkEndingZeroes(number, zeroes) {
		fmt.Println("Giusto! ", number, " ha ", zeroes, " zeri!")
	} else {
		fmt.Println("Sbagliato! ", number, " NON ha ", zeroes, " zeri!")
	}
}

func checkEndingZeroes(number int, zeroes int) bool {
	return number%powInt(10, zeroes) == 0
}

func powInt(base int, exponent int) int {
	if exponent < 1 {
		return 1
	} else {
		return powInt(base, exponent-1) * base
	}
}
