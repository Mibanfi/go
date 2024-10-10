package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var radius int
	if len(os.Args) > 1 {
		radius, _ = strconv.Atoi(os.Args[1])
	} else {
		radius = 0
	}
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
		fmt.Print("So your number is ", radius, ". Good. Please remember to input an argument next time.\n\n")
	}
	var radiusFloat float64
	radiusFloat = float64(radius)
	// First half
	for i := 0; i < radius; i++ {
		for j := 0; j < int((radiusFloat-math.Sqrt(math.Pow(radiusFloat, 2)-math.Pow(radiusFloat-float64(i), 2)))+0.5); j++ {
			fmt.Print("  ")
		}
		for h := 0; h < int((2*math.Sqrt(math.Pow(radiusFloat, 2)-math.Pow(radiusFloat-float64(i), 2)))+0.5); h++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
	// Second half
	for i := radius; i > 0; i = i - 1 {
		for j := 0; j < int((radiusFloat-math.Sqrt(math.Pow(radiusFloat, 2)-math.Pow(radiusFloat-float64(i), 2)))+0.5); j++ {
			fmt.Print("  ")
		}
		for h := 0; h < int((2*math.Sqrt(math.Pow(radiusFloat, 2)-math.Pow(radiusFloat-float64(i), 2)))+0.5); h++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
	fmt.Println("")
}
