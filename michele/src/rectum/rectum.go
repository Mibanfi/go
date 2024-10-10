package main

import (
	"fmt"
)

func main() {
	var m1, m2, q1, q2 int
	var x, y int
	fmt.Println("Inserisci m della prima retta:")
	fmt.Scan(&m1)
	fmt.Println("Inserisci q della prima retta:")
	fmt.Scan(&q1)
	fmt.Println("Inserisci m della seconda retta:")
	fmt.Scan(&m2)
	fmt.Println("Inserisci q della seconda retta:")
	fmt.Scan(&q2)
	/*
		y = m1 * x + q1
		y = m2 * x + q2

		m2 * x + q2 = m1 * x + q1
		m2 * x - m1 * x = q1 - q2
		x * (m2 - m1) = q1 - q2
		x = (q1 - q2)/(m2 - m1)
	*/
	if m1 == m2 {
		fmt.Println("Le rette sono parallele!")
	} else {
		x = (q1 - q2) / (m2 - m1)
		y = m1*x + q1
		fmt.Println("Intersezione in (", x, ",", y, ")")
	}
}
