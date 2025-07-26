package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var wg = sync.WaitGroup{}

func main1() {
	// go routines

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total execution tine: %v\n", time.Since(t0))

}

func dbCall(i int) {
	//simulate db call delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result from the database is: %v\n", dbData[i])
	wg.Done()
}
