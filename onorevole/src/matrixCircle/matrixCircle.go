package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// Initializing
	var radius int
	var symbol string
	var vPos, hPos int

	if len(os.Args) >= 3 {
		symbol = os.Args[1]
		radius, _ = strconv.Atoi(os.Args[2])
	} else {
		// Taking input
		fmt.Print("Insert radius: ")
		fmt.Scan(&radius)
		radius++
		fmt.Print("Choose symbol: ")
		fmt.Scan(&symbol)
		fmt.Print("\n")
	}

	// Making the 2D slice which will be used as a canvas
	var canvas [][]bool
	canvas = make([][]bool, radius*2)
	for i := 0; i < len(canvas); i++ {
		canvas[i] = make([]bool, radius*2)
	}

	// Setting up one quarter
	for vPos = radius; vPos > 0; vPos-- {
		for hPos = radius; hPos > radius-(chordLength(radius, radius-vPos)/2); hPos-- {
			canvas[hPos][vPos] = true
		}
	}

	// Closing up: mirroring the other quarters
	for vPos = 0; vPos < radius; vPos++ {
		for hPos := 0; hPos < radius; hPos++ {
			canvas[radius*2-1-hPos][vPos] = canvas[hPos][vPos]
			canvas[hPos][radius*2-1-vPos] = canvas[hPos][vPos]
			canvas[radius*2-1-hPos][radius*2-1-vPos] = canvas[hPos][vPos]
		}
	}

	//Drawing everything
	for vPos = 0; vPos < radius*2; vPos++ {
		for hPos = 0; hPos < radius*2; hPos++ {
			if canvas[hPos][vPos] {
				fmt.Print(symbol, " ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func chordLength(radius int, distance int) int {
	var radiusF, distanceF, resultF float64
	radiusF = float64(radius)
	distanceF = float64(distance)
	resultF = 2 * math.Sqrt(math.Pow(radiusF, 2)-math.Pow(distanceF, 2))
	return int(resultF + 0.5)
}
