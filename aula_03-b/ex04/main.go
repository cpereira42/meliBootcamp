package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	variavel := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)

	fmt.Println("Inseção")
	orderInse(variavel)
	orderInse(variavel2)
	orderInse(variavel3)

	fmt.Println("\nSort")
	ordeSort(variavel)
	ordeSort(variavel2)
	ordeSort(variavel3)

	fmt.Println("\nSeleçao")
	orderSel(variavel)
	orderSel(variavel2)
	orderSel(variavel3)

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
