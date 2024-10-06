package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var radius int
	radius, _ = strconv.Atoi(os.Args[1])
	if radius == 0 {
		fmt.Println("It seems like you didn't type any argument to this program, or you typed 0. This program needs an argument other than 0, to use as the circle's radius. Now, input the number. What should be the circle's radius?")
		fmt.Scan(&radius)
		if radius == 0 {
			fmt.Println("I think you didn't understand. You need to type an integer number, which will be the circle's radius.")
			fmt.Scan(&radius)
			if radius == 0 {
				fmt.Println("Ugh, listen, I gave you a chance. I don't know if you're messing with me or what, but you know what? You can go f%$ck yourself. I don't get paid enough to deal with idiots like you.")
				return
			}
		}
		fmt.Print(radius, ". Good. Please remember to input an argument next time.\n\n")
	}
	for i := 0; i < radius; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("  ")
		}
		fmt.Println("* ")
	}
}
