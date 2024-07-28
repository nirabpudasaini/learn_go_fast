package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("Total time: %v\n", time.Since(t0))
	fmt.Printf("Results: %v\n", results)
}

func dbCall(i int) {
	//Simulate Db vall delay
	var delay float32 = rand.Float32() * 1000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Data from db is:", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("current result: %v \n", results)
	m.RUnlock()
}
