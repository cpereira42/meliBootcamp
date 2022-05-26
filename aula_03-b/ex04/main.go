package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	fmt.Println("inicia")
	go checkInse(c1)
	<-c1
	go checkSort(c2, c1)
	<-c2
	go checkSel(c3, c2)
	<-c3
}
func checkInse(out chan<- int) {
	fmt.Println("\nInseção")
	orderInse(rand.Perm(100))
	orderInse(rand.Perm(1000))
	orderInse(rand.Perm(10000))
	out <- 1
	close(out)
}
func checkSort(out chan<- int, in <-chan int) {
	fmt.Println("\nSort")
	ordeSort(rand.Perm(100))
	ordeSort(rand.Perm(1000))
	ordeSort(rand.Perm(10000))
	out <- 1
	close(out)
}

func checkSel(out chan<- int, in <-chan int) {
	fmt.Println("\nSeleçao")
	orderSel(rand.Perm(100))
	orderSel(rand.Perm(1000))
	orderSel(rand.Perm(10000))
	out <- 1
	close(out)
}

func ordeSort(variavel []int) {
	ti := time.Now()
	sort.Ints(variavel)
	tf := time.Now()
	fmt.Printf("%s\n", tf.Sub(ti))
}

func orderSel(variavel []int) {

	ti := time.Now()
	var n = len(variavel)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if variavel[j] < variavel[minIdx] {
				minIdx = j
			}
		}
		variavel[i], variavel[minIdx] = variavel[minIdx], variavel[i]
	}
	tf := time.Now()
	fmt.Printf("%s\n", tf.Sub(ti))

}

func orderInse(variavel []int) {

	ti := time.Now()
	var n = len(variavel)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if variavel[j-1] > variavel[j] {
				variavel[j-1], variavel[j] = variavel[j], variavel[j-1]
			}
			j = j - 1
		}
	}
	tf := time.Now()
	fmt.Printf("%s\n", tf.Sub(ti))
}

func printArray(variavel []int) {
	for _, num := range variavel {
		fmt.Print(num, ",")
	}
	fmt.Print("\n")
}
