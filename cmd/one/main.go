package main

import "fmt"

func main() {
	//AArays and slices
	intArr := [...]int{2, 3, 4}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("the lenght is %v and the capacity is %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Printf("the lenght is %v and the capacity is %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{7, 8}
	intSlice = append(intSlice, intSlice2...)

	println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	println(intSlice3)

	//Maps

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Printf("%v\n", myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "saraf": 54}
	fmt.Println(myMap2["json"]) //0

	var age, ok = myMap2["Adam"]

	if ok {
		fmt.Printf("age %v", age)
	} else {
		fmt.Println("Invalid")
	}

	//iteration

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}
}
