package main
import "fmt"
import "math"

//Objective: create a program that converts an integer number (e.g. 1234) into its equivalent in words (e.g. one thousand two hundred thirty-four).
//Limitations: may not convert an integer directly into a string (e.g. 1234 -> "1234")

func main() {
	var num,numArr int
	
	//Conversion tables
	var hundreds = [10]string{"","One Hundred ","Two Hundred ","Three Hundred ","Four Hundred ","Five Hundred ","Six Hundred ","Seven Hundred ","Eight Hundred ","Nine Hundred "}
	var tens = [10]string{"","","Twenty","Thirty","Fourty","Fifty","Sixty","Seventy","Eighty","Ninety"}
	var units = [10]string{"","One ","Two ","Three ","Four ","Five ","Six ","Seven ","Eight ","Nine "}
	var teens = [10]string{"Ten ","Eleven ","Twelve ","Thirteen ","Fourteen ","Fifteen ","Sixteen ","Seventeen ","Eighteen ","Nineteen "}
	var magnitude = [8]string{"","Thousand ","Million ","Billion ","Trillion ","Quadrillion ","Pentillion ","Sextillion "}
	
	//Ask for input
	fmt.Print("Input an integer number: ")
	fmt.Scan(&num)
	
	//The two main arrays
	var numarr[digCount] int
	numArr = intToArr(num)
	var wordArr[len.numArr] string

	//Setting up the for cycle
	var i int
	var word string
	for i=0; i<len.numArr; {
		//Units
		if wordArr[i+1] == 1 {
			wordArr[i] = teens[numArr[i]]
		} else {
			wordArr[i] = units[numArr[i]]
		}
		//Magnitude
		wordArr[i] = wordArr[i]+magnitude[i/3]
		i++
		//Tens
		wordArr[i] = tens[numArr[i]]
		if numArr[i-1]!=0 {
			wordArr[i]=wordArr[i]+"-"
		} else if numArr[i] != 0 {
			wordArr[i]=wordArr[i]+" "
		}
		i++
		//Hundreds
		wordArr[i] = hundreds[numArr[i]]
		i++
	}
}

func intToArr(num int) {
	//Calculate number of digits
	var digCount int
	digCount = digCount(num)
	//Create array to store digits
	var dig [digCount]int
	//Store digits into array
	for currentDig:=0; currentDig<digCount; currentDig++ {
		dig[currentDig] = (num%powInt(10,currentDig+1))/powInt(10,currentDig)
	}
	return dig
}

func digCount(n int) {
	var count, result int
	for count=0; result==0; count++ {
		result=(n)/powInt(10,count)
	}
	return count
}

func powInt(base int, exponent int) {
	if exponent == 0 {
		return base
	} else {return powInt(base, exponent-1)*base}
}
