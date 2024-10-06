package main
import (
	"fmt"
	"strings"
)
func main() {
	var n,i int
	fmt.Print("Quanto devono essere lunghi i lati del triangolo? ")
	fmt.Scan(&n)
	fmt.Println("")
	if n == 0 {
		fmt.Println("Che pirla che sei")
	} else { 
		fmt.Println("*")
	}
	for i=1; i<n-1; i++ {
		fmt.Println("* "+strings.Repeat("  ",i-1)+"*")
	}
	if n > 1 {
		fmt.Println(strings.Repeat("* ",n))
	} else if n == 1 {
		fmt.Println("Deludente? Senti, provaci tu a disegnare un triangolo 1x1 in un terminale")
	}
	fmt.Println("")
}
