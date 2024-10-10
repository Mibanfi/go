package main
import "fmt"
func main() {
	var n,i,j int
	fmt.Print("Quanto devono essere lunghi i lati del triangolo?")
	fmt.Scan(&n)
	fmt.Println("*")
	for i=1; i<n; i++ {
		fmt.Print("*")
		for j=0; j<i; j++ {
			fmt.Print("  ")
		}
		fmt.Println("*")	
	}
	for i=0; i<n; i++ {
		fmt.Print("* ")
	}
	fmt.Println("*")
}
