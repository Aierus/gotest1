package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func Problem1(x int) {
	// Time duration code taken from https://pkg.go.dev/time#Duration
	t0 := time.Now()
	exampleArray := [1000]int{}
	t1 := time.Now()
	fmt.Printf("Initializing an array of len %d took %v miliseconds\n", len(exampleArray), t1.Sub(t0))

	t0 = time.Now()
	exampleSlice := make([]int, x)
	t1 = time.Now()
	fmt.Printf("Initializing an slice of len %d took %v miliseconds\n", len(exampleSlice), t1.Sub(t0))

	t0 = time.Now()
	exampleMap := make(map[int]int)
	t1 = time.Now()
	fmt.Printf("Initializing an map of len %d took %v miliseconds\n", len(exampleMap), t1.Sub(t0))

	for i := 0; i < x; i++ {
		if i == 1 {
			t0 := time.Now()
			exampleArray[i] = i
			t1 := time.Now()
			fmt.Printf("incrementing an array by 1 took %d miliseconds\n", t1.Sub(t0))
		}
		exampleArray[i] = i
	}
	for i, _ := range exampleSlice {
		if i == 1 {
			t0 := time.Now()
			exampleSlice[i] = i
			t1 := time.Now()
			fmt.Printf("incrementing a slice by 1 took %d miliseconds\n", t1.Sub(t0))
		}
		exampleSlice[i] = i
	}
	for i, v := range exampleMap {
		if i == 1 {
			t0 := time.Now()
			exampleMap[i] = i
			t1 := time.Now()
			fmt.Printf("incrementing a map by 1 took %d miliseconds\n", t1.Sub(t0))
		}
		exampleMap[i] = i
		exampleMap[v] = v
	}
}

type Slice []int
type Array [100]int

func (a Slice) Len() int           { return len(a) }
func (a Slice) Less(i, j int) bool { return a[i] < a[j] }
func (a Slice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a Array) Len() int           { return len(a) }
func (a Array) Less(i, j int) bool { return a[i] < a[j] }
func (a Array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Problem2(x int) {
	exSlice := make([]int, x)
	exArray := [100]int{}

	for i, _ := range exSlice {
		exSlice[i] = rand.Int()
	}
	for i, _ := range exArray {
		exArray[i] = rand.Int()
	}

	t0 := time.Now()
	sort.Sort(Slice(exSlice))
	t1 := time.Now()
	fmt.Printf("sorting a slice using sort.Sort took %d miliseconds\n", t1.Sub(t0))

	t0 = time.Now()
	sort.Stable(Slice(exSlice))
	t1 = time.Now()
	fmt.Printf("sorting a slice using sort.Stable took %d miliseconds\n", t1.Sub(t0))

	t0 = time.Now()
	sort.Sort(Array(exArray))
	t1 = time.Now()
	fmt.Printf("sorting an array using sort.Sort took %d miliseconds\n", t1.Sub(t0))

	t0 = time.Now()
	sort.Stable(Array(exArray))
	t1 = time.Now()
	fmt.Printf("sorting an array using sort.Stable took %d miliseconds\n", t1.Sub(t0))
}

func main() {
	corrections := os.Args[1]
	fmt.Printf("Input from cmd line is %s\n", corrections)
	x, _ := strconv.Atoi(corrections)
	if x > 1000 {
		x = 1000
	}
	Problem1(x)
	Problem2(x)
}
