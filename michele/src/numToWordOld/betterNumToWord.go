package main

import "fmt"

func main() {
}

func digitSlice(num int) {
	var digitSlice []int = make([]int, countDigit(num))
	var index int
	for index = 0; index < countDigit(num); index++ {
		digitSlice[index] = (num/powInt(10,index))%powInt(10,index+1)
	}
	return digitSlice
}

func countDigit(num int) {
	if num/10 == 0 {
		return 0
	} else {
		return countDigit(num/10)+1
	}
}

func powInt(base int, exponent int) {
	if exponent == 0 {
		return 1
	} else {
		return (powInt(base, exponent-1)*base)
	}
}
