package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

// Q1 Corrections, only graded for slice and map

func Problem1(x int) {
	// Time duration code taken from https://pkg.go.dev/time#Duration
	t0 := time.Now()
	exampleArray := [1000]int{}
	t1 := time.Now()
	fmt.Printf("Initializing an array of len %d took %v\n", len(exampleArray), t1.Sub(t0))

	t0 = time.Now()
	exampleArray100 := [100]int{}
	t1 = time.Now()
	fmt.Printf("Initializing an array of len %d took %v\n", len(exampleArray100), t1.Sub(t0))

	t0 = time.Now()
	exampleArray10 := [10]int{}
	t1 = time.Now()
	fmt.Printf("Initializing an array of len %d took %v\n", len(exampleArray10), t1.Sub(t0))

	t0 = time.Now()
	exampleSlice := make([]int, x)
	t1 = time.Now()
	fmt.Printf("Initializing an slice of len %d took %v\n", len(exampleSlice), t1.Sub(t0))

	t0 = time.Now()
	exampleMap := make(map[int]int)
	t1 = time.Now()
	fmt.Printf("Initializing an map of len %d took %v\n", len(exampleMap), t1.Sub(t0))
	fmt.Printf("The map is dynamic, and thus has not been assigned a length yet\n")

	for i := 0; i < 1000; i++ {
		if i == 1 {
			t0 := time.Now()
			exampleArray[i] = i
			t1 := time.Now()
			fmt.Printf("incrementing an array by 1 took %dns\n", t1.Sub(t0))
		}
		exampleArray[i] = i
	}
	// fmt.Printf("Example array of length 100 after assignment == ")
	// fmt.Println(exampleArray100)

	for i, _ := range exampleSlice {
		if i == 1 {
			t0 := time.Now()
			exampleSlice[i] = i
			t1 := time.Now()
			fmt.Printf("incrementing a slice by 1 took %dns\n", t1.Sub(t0))
		}
		exampleSlice[i] = i
	}
	// fmt.Printf("Example slice of length %d after assignment == ", len(exampleSlice))
	// fmt.Println(exampleSlice)
	for i := 0; i < x; i++ {
		v := i
		if i == 1 {
			t0 := time.Now()
			exampleMap[i] = v
			t1 := time.Now()
			fmt.Printf("incrementing a map by 1 took %dns\n", t1.Sub(t0))
		}
		exampleMap[i] = v
	}
	// fmt.Printf("Example map of length %d after assignment == ", len(exampleMap))
	// fmt.Println(exampleMap)
}

type Interface []int

func (a Interface) Len() int           { return len(a) }
func (a Interface) Less(i, j int) bool { return a[i] < a[j] }
func (a Interface) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Problem2(x int) {
	exSlice := make([]int, x)
	exArray1000 := [1000]int{}
	exArray100 := [100]int{}

	for i, _ := range exSlice {
		exSlice[i] = rand.Int()
	}

	for i, _ := range exArray1000 {
		exArray1000[i] = rand.Int()
	}

	for i, _ := range exArray100 {
		exArray100[i] = rand.Int()
	}

	t0 := time.Now()
	sort.Sort(Interface(exSlice))
	t1 := time.Now()
	fmt.Printf("sorting a slice of length %d using sort.Sort took %dns\n", len(exSlice), t1.Sub(t0))

	t0 = time.Now()
	sort.Stable(Interface(exSlice))
	t1 = time.Now()
	fmt.Printf("sorting a slice using sort.Stable took %dns\n", t1.Sub(t0))

	t0 = time.Now()
	sort.Sort(Interface(exArray1000[:])) // Convert array to slice using [:]
	t1 = time.Now()
	fmt.Printf("sorting an array of length %d using sort.Sort took %dns\n", len(exArray1000), t1.Sub(t0))

	t0 = time.Now()
	sort.Stable(Interface(exArray1000[:]))
	t1 = time.Now()
	fmt.Printf("sorting an array of length %d using sort.Stable took %dns\n", len(exArray1000), t1.Sub(t0))

	t0 = time.Now()
	sort.Sort(Interface(exArray100[:])) // Convert array to slice using [:]
	t1 = time.Now()
	fmt.Printf("sorting an array of length %d using sort.Sort took %dns\n", len(exArray100), t1.Sub(t0))

	t0 = time.Now()
	sort.Stable(Interface(exArray100[:]))
	t1 = time.Now()
	fmt.Printf("sorting an array of length %d using sort.Stable took %dns\n", len(exArray100), t1.Sub(t0))

}

func main() {
	// Get argument `x` from command line, first arg after program name
	x := os.Args[1]
	// Echo argument `x` from command line to ensure correct input
	fmt.Printf("Input from cmd line is %s\n", x)
	// Convert command line argument `x` to int `y` for use in program
	y, _ := strconv.Atoi(x)
	Problem1(y)
	Problem2(y)
}
