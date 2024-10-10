package main

import (
	"fmt"
	"strings"
)

func main() {
	var n, i, j, c int
	c = 1
	i = 0
	fmt.Println("")
	fmt.Print("Quanto deve essere lungo il diametro del simbolo? (Minimo 5, solo numeri dispari) ")
	fmt.Scan(&n)
	if n < 0 {
		n = n * (-1)
		c = -1
		i = 4
	}
	if n < 5 || (n/2)*2 == n {
		fmt.Println("Ma sei deficiente?")
	} else {
		fmt.Println("")
		for i = i; i < 5 && i > -1; i = i + c {
			switch i {
			case 0:
				fmt.Println("* " + strings.Repeat("  ", n/2-1) + strings.Repeat("* ", n/2+1))
			case 1:
				for j = 0; j < n/2-1; j++ {
					fmt.Println("* " + strings.Repeat("  ", n/2-1) + "* " + strings.Repeat(" ", n/2))
				}
			case 2:
				fmt.Println(strings.Repeat("* ", n))
			case 3:
				for j = 0; j < n/2-1; j++ {
					fmt.Println(strings.Repeat("  ", n/2) + "* " + strings.Repeat("  ", n/2-1) + "* ")
				}
			case 4:
				fmt.Println(strings.Repeat("* ", n/2+1) + strings.Repeat("  ", n/2-1) + "* ")
			}
		}
		fmt.Println("")
	}
}
