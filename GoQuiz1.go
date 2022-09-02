package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Problem1(x int) {
	// Time duration code taken from https://pkg.go.dev/time#Duration
	t0 := time.Now()
	exampleArray := [...]int{}
	t1 := time.Now()
	fmt.Printf("Initializing an array of len %d took %v seconds\n", x, t1.Sub(t0))

	t0 = time.Now()
	exampleSlice := make([]int, x)
	t1 = time.Now()
	fmt.Printf("Initializing an slice of len %d took %v seconds\n", x, t1.Sub(t0))

	t0 = time.Now()
	exampleMap := make(map[int]int)
	t1 = time.Now()
	fmt.Printf("Initializing an map of len %d took %v seconds\n", x, t1.Sub(t0))

	for i := 0; i < x; i++ {
		if i == 1 {
			t0 := time.Now()
			exampleArray[i] = i
			t1 := time.Now()
			fmt.Printf("incrementing an array by 1 took %d seconds", t1.Sub(t0))
		}
		exampleArray[i] = i
	}
	for i, _ := range exampleSlice {
		if i == 1 {
			t0 := time.Now()
			exampleSlice[i] = i
			t1 := time.Now()
			fmt.Printf("incrementing a slice by 1 took %d seconds", t1.Sub(t0))
		}
		exampleSlice[i] = i
	}
	for i, v := range exampleMap {
		if i == 1 {
			t0 := time.Now()
			exampleMap[i] = i
			t1 := time.Now()
			fmt.Printf("incrementing a map by 1 took %d seconds", t1.Sub(t0))
		}
		exampleMap[i] = i
		exampleMap[v] = v
	}
}

func Problem2(x int) {
	exSlice := make([]int, x)
	exArray := [...]int{}
	for i, _ := range exSlice {
		exSlice[i] = rand.Int()
	}
	for i, _ := range exArray {
		exArray[i] = rand.Int()
	}

	// Sort.interface example from https://yourbasic.org/golang/how-to-sort-in-go/
	func (a []int) Len() int           { return len(a) }
	func (a []int) Less(i, j int) bool { return a[i] < a[j] }
	func (a []int) Swap(i, j int)      { a[i] = a[j] }

	t0 := time.Now()
	sort.Sort(exSlice)
	t1 := time.Now()
	fmt.Printf("sorting a slice using sort.Sort took %d seconds", t1.Sub(t0))
	

	t0 = time.Now()
	sort.Stable(exSlice)
	t1 = time.Now()
	fmt.Printf("sorting a slice using sort.Stable took %d seconds", t1.Sub(t0))


	t0 = time.Now()
	sort.Sort(exArray)
	t1 = time.Now()
	fmt.Printf("sorting an array using sort.Sort took %d seconds", t1.Sub(t0))
	
	t0 = time.Now()
	sort.Stable(exArray)
	t1 = time.Now()
	fmt.Printf("sorting an array using sort.Stable took %d seconds", t1.Sub(t0))
}

func main() {
	programName := os.Args[0]
	x := strconv.Atoi(programName)
	Problem1(x)
	Problem2(x)
}
