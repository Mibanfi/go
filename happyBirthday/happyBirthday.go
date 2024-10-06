package plutopos

import "fmt"

func main() {
	var age, i uint64
	fmt.Print("Your age is ")
	fmt.Scan(&age)
	for i = 1; i < age; i++ {
		fmt.Println(numToWord(i))
	}
	fmt.Println("..aaaaaaand...")
	fmt.Print(numToWord(age), "! HAPPY BIRTHDAY!!")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("I love you :D")
}
