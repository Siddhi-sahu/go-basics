package main

import "fmt"

import {
	"log"
	"os"
}

func main() {
	//declaring variables

	// a := "heeyy"
 
	// fmt.Println(a)

	// var f string
	// f = "withhh"

	// fmt.Println(f)

	//Error handling
	f, err := os.Open("filename.txt")

	if err != nil {
		log.Fatal(err)
	}

	//Do something with file

	n, err := f.Read()

	if err != nil{

	}

	//Arrays and Slices

	//arrays are static while slices are dynamic

	prime := [6]int{2, 3, 4, 7, 11, 13}

	var s []int = prime[1:4]

	fmt.Println(s)

	//Maps
	m := make(map[string]int)

	m["ONE"] = 1

	fmt.Println("map", m)

	// for range loop over slices

	k := []int{1, 2, 3, 4, 5}

	for i, v := range k {
		fmt.Println(i, v)
	}
}
