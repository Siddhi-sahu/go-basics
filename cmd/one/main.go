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
	var result, remainder, err = intDivision(numerator, denominator)

	if err != nil {
		fmt.Printf("%v\n", err)
	} else if remainder == 0 {
		fmt.Printf("the result of integer division is %v", result)

	} else {
		fmt.Printf("the result of integer division is %v and remainder is %v", result, remainder)

	}

}

func PrintMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}

//dataTypes
// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	var intNum int16 = 3278
// 	fmt.Println(intNum)
// 	fmt.Println(len("γ"))                    //number of bytes in string
// 	fmt.Println(utf8.RuneCountInString("γ")) //actual length of string

// 	var myString = "text"
// 	myString2 := "zz"

// 	fmt.Println(myString, myString2)

// }
