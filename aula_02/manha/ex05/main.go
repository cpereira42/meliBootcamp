package main

import "fmt"

const (
	Dog = "dog"
	Cat = "cat"
	Tarantula = "tarantula"
	Haminster = "haminster"
)

func main(){


	animaldog, msg := Animal(Dog)
	if msg != ""{
		fmt.Println("Erro :",msg)
	}
	animalCat, msg := Animal(Cat)
	animalHamister, msg := Animal(Haminster)
	animalTarantula, msg := Animal(Tarantula)

	if msg != ""{
		fmt.Println("Erro :",msg)
	}

	var amount float64
	amount+= animaldog(0)
	amount+= animalCat(0)
	amount+= animalHamister(1)
	amount+= animalTarantula(1)

	fmt.Println("Irá precisar de",amount)

}

func Animal(tipo string) (func(qtt int)(float64),string) {
	switch tipo {
	case Dog:
		return calcDog, ""
	case Cat:
		return calcCat, ""
	case Haminster:
		return calcHamister, "s"
	case Tarantula:
		return calcTarantula, ""
	default:
		return invalido,"Animal não cadastrado"
	}
}

func invalido(qtt int) float64{
	return float64(0)
}
func calcDog(qtt int) float64{
	return float64(qtt * 10)
}

func calcCat(qtt int) float64{
	return float64(qtt * 5)
}

func calcHamister(qtt int)float64{
	return float64(qtt) * 0.250
}

func calcTarantula(qtt int)float64{
	return  float64(qtt) * 0.150
}