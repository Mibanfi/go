package plutopos

import "fmt"

func NumToWord(number uint64) string {
	// Matrix used for conversion
	var conversionTable [3][10]string

	// Units
	conversionTable[0][0] = ""
	conversionTable[0][1] = "One"
	conversionTable[0][2] = "Two"
	conversionTable[0][3] = "Three"
	conversionTable[0][4] = "Four"
	conversionTable[0][5] = "Five"
	conversionTable[0][6] = "Six"
	conversionTable[0][7] = "Seven"
	conversionTable[0][8] = "Eight"
	conversionTable[0][9] = "Nine"

	// Tens
	conversionTable[1][0] = ""
	conversionTable[1][1] = "Ten"
	conversionTable[1][2] = "Twenty"
	conversionTable[1][3] = "Thirty"
	conversionTable[1][4] = "Fourty"
	conversionTable[1][5] = "Fifty"
	conversionTable[1][6] = "Sixty"
	conversionTable[1][7] = "Seventy"
	conversionTable[1][8] = "Eighty"
	conversionTable[1][9] = "Ninety"

	// Hundreds
	conversionTable[2][0] = ""
	conversionTable[2][1] = "One Hundred"
	conversionTable[2][2] = "Two Hundred"
	conversionTable[2][3] = "Three Hundred"
	conversionTable[2][4] = "Four Hundred"
	conversionTable[2][5] = "Five Hundred"
	conversionTable[2][6] = "Six Hundred"
	conversionTable[2][7] = "Seven Hundred"
	conversionTable[2][8] = "Eight Hundred"
	conversionTable[2][9] = "Nine Hundred"

	// Order of magnitude
	var magnitudeTable [7]string
	magnitudeTable[0] = ""
	magnitudeTable[1] = "Thousand"
	magnitudeTable[2] = "Million"
	magnitudeTable[3] = "Billion"
	magnitudeTable[4] = "Trillion"
	magnitudeTable[5] = "Quadrillion"
	magnitudeTable[6] = "Quintillion"

	// Teens
	var teensTable [10]string
	teensTable[0] = "Ten"
	teensTable[1] = "Eleven"
	teensTable[2] = "Twelve"
	teensTable[3] = "Thirteen"
	teensTable[4] = "Fourteen"
	teensTable[5] = "Fifteen"
	teensTable[6] = "Sixteen"
	teensTable[7] = "Seventeen"
	teensTable[8] = "Eighteen"
	teensTable[9] = "Nineteen"

	//Actual code
	if number == 0 {
		return "Zero"
	}
	var digits [20]int
	digits = digitsToArray(number)
	var word string
	fmt.Println("Your chosen number was:", number)
	for i := DigitCount(number) - 1; i >= 0; i = i - 1 {
		// Checks for "teen"-type numbers
		if digits[i] == 1 && i%3 == 1 {
			word = word + teensTable[digits[i-1]]
			i = i - 1
		} else {
			word = word + conversionTable[i%3][digits[i]]
		}
		// Adds the word for the current digit, using the conversion table
		/* Adds "and" if near the end
		if i == 2 && digits[i-1]+digits[i-2] > 0 {
			word = word + " and"
		} */
		// Adds a space, a dash or nothing, depending on what comes next
		if i%3 == 1 && digits[i-1] > 0 && digits[i] > 1 {
			word = word + "-"
		} else if digits[i] > 0 {
			word = word + " "
		}
		// Adds order of magnitude, if necessary
		if i%3 == 0 && digits[i]+digits[i+1]+digits[i+2] > 0 {
			word = word + magnitudeTable[i/3] + " "
		}

	}
	return word
}

func DigitCount(num uint64) int {
	if num == 0 {
		return 0
	} else {
		return DigitCount(num/10) + 1
	}
}

func digitsToArray(num uint64) [20]int {
	var digits [20]int
	var count int
	count = DigitCount(num)
	for i := 0; i < count; i++ {
		digits[i] = int((num / uint64(PowInt(10, i))) % 10)
	}
	return digits
}

func PowInt(base int, exponent int) int {
	if exponent < 1 {
		return 1
	} else {
		return PowInt(base, exponent-1) * base
	}
}
