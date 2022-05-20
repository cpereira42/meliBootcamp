package main

import "fmt"

func main(){

	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println("Sua empresa tem",len(employees))
	limite := 21
	sum := 0

	for key, element := range employees{
		if element > 21{
			fmt.Print(key,",")
			sum ++
		}
	}
	if (sum > 0){
		fmt.Println("são maiores de",limite,"anos")
	}
		
	fmt.Println("Adicionando Frederico")
	employees["Frederico"] = 25
	fmt.Println("Excluíndo Pedro")
	delete(employees, "Pedro")
}
