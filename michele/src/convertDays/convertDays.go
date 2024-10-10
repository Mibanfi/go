package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Initialize
	const year = 365
	const month = 30
	const week = 7
	var days int
	var err error

	// Take input
	if len(os.Args) == 2 {
		days, err = strconv.Atoi(os.Args[1])
	} else {
		fmt.Println("Please insert a number of days.")
		_, err = fmt.Scan(&days)
	}

	// Handle errors
	for err != nil {
		fmt.Println("Error:", err, "\nPlease insert a new number of days.")
		_, err = fmt.Scan(&days)
	}

	// Print results
	fmt.Println(days, "days =", days/year, "years,", (days%year)/month, "months,", (days%month)/week, "weeks,", (days % week), "days.")
}
