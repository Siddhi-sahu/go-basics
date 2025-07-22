package main

import (
	"errors"
	"fmt"
)

func main() {
	// var printValue string = "helloo"
	// PrintMe(printValue)

	var numerator int = 10
	var denominator int = 2
	var result, remainder int = intDivision(numerator, denominator)

	fmt.Printf("the result of integer division is %v and remainder is %v", result, remainder)
}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by 0")
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder
}

//dataTypes
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 3278
	fmt.Println(intNum)
	fmt.Println(len("γ"))                    //number of bytes in string
	fmt.Println(utf8.RuneCountInString("γ")) //actual length of string

	var myString = "text"
	myString2 := "zz"

	fmt.Println(myString, myString2)

}
