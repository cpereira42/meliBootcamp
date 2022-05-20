package main

import "fmt"

func main(){

	fmt.Printf("A média é %2f",calcMedia(5.3,5.4,6))
}

func calcMedia(valores ...float32) float32 {

	var sum float32
	for i:=0; i< len(valores); i++{
		sum +=valores[i]
	}
	return sum / float32(len(valores))
	
}