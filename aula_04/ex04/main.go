package main

import (
	"errors"
	"fmt"
)

func main() {

	msg, err := calcSalarioNew(50, 30.0)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		e1 := fmt.Errorf("error: %w", err)
		fmt.Println(e1)
		fmt.Println("error:", errors.Unwrap(e1))

	} else {
		fmt.Println(msg)
	}
}

func calcSalarioNew(horasTrabalhadas int, valorHora float64) (float64, error) {
	if horasTrabalhadas < 80 {
		return 0, errors.New("o trabalhador não pode ter trabalhado menos de 80 horas por mês")
	}
	salario := float64(horasTrabalhadas) * valorHora
	if salario > 20000 {
		salario *= 0.9
	}
	return float64(salario), nil
}
