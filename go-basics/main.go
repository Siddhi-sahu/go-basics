package main

import "fmt"

func main() {
	// fmt.Println("this is itt")
	// anotherFile()

	sum := 0

	for sum < 1000 {
		sum += sum
		sum = sum + 1
	}

	fmt.Println(sum)
}
