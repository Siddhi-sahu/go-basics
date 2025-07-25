package main

import (
	"fmt"
	"time"
)

func main() {
	//Pointers

	var p *int32 = new(int32)
	var i int32

	fmt.Printf("the value p points at: %v\n", *p)
	fmt.Printf("the value of i: %v\n", i)

	p = &i
	*p = 10

	fmt.Printf("the value p points at: %v\n", *p)
	fmt.Printf("the value of i: %v\n", i)
	//structs and interfaces

	var myEngine gasEngine = gasEngine{20, 40}
	fmt.Println(myEngine.gallons, myEngine.mpg)
	fmt.Printf("total miles left int the tank: %v\n", myEngine.milesLeft())

	// preset capactity
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without prelocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with prelocation: %v\n", timeLoop(testSlice2, n))

	//Strings, Runes and bytes

	var myString string = "résumé"
	var indexed = myString[0]

	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

}

type gasEngine struct {
	mpg     uint16
	gallons uint16
}

func (e gasEngine) milesLeft() uint16 {
	return e.gallons * e.mpg

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()

	for len(slice) < n {
		slice = append(slice, 1)
	}

	return time.Since(t0)

}
