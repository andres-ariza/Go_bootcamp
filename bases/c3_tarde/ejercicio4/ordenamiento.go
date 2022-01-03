package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sorts
func insertionSort(items []int, c chan time.Duration, s chan string) {
	start := time.Now()
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
	s <- "insertion"
	c <- time.Since(start)
}
func bubbleSort(items []int, c chan time.Duration, s chan string) {
	start := time.Now()
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	s <- "bubble"
	c <- time.Since(start)

}
func selectionSort(items []int, c chan time.Duration, s chan string) {
	start := time.Now()
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	s <- "selection"
	c <- time.Since(start)
}

func main() {
	v1 := rand.Perm(100)
	v2 := rand.Perm(1000)
	v3 := rand.Perm(10000)
	c, s := make(chan time.Duration), make(chan string)

	// 100 round
	fmt.Println("----- 100 round -----")
	go insertionSort(v1, c, s)
	go bubbleSort(v1, c, s)
	go selectionSort(v1, c, s)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)

	// 1000 round
	fmt.Println("----- 1000 round -----")
	go insertionSort(v2, c, s)
	go bubbleSort(v2, c, s)
	go selectionSort(v2, c, s)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)

	// 10000 round
	fmt.Println("----- 10000 round -----")
	go insertionSort(v3, c, s)
	go bubbleSort(v3, c, s)
	go selectionSort(v3, c, s)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)
	fmt.Printf("Termino el %s sort en %s\n", <-s, <-c)
}
