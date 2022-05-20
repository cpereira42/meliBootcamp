package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
	)
	
func main(){

	minhaFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)
	defaultFunc, err := operation("default")

	fmt.Println(minhaFunc, err)
	

	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	defaultValue := defaultFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Printf("Mean: %.2f Min: %.2f Max: %.2f default: %.2f\n",
	averageValue, minValue, maxValue, defaultValue)

}

func operation (operator string) (func(valores ...int)(float32), error){
	switch operator {
	case average:
		return averageFunc, nil
	case minimum:
		return minhaFunc, nil
	case maximum:
		return maxFunc, nil
	default:
		return defaultFunc, errors.New("operador não informado")
	}
	
}

func minhaFunc(valores ...int) (float32){

	min:= valores[0]
	for i:= 1; i< len(valores); i++{
		if (valores[i] < min){
			min = valores[i]
		}
	}
	return float32(min)
}

func averageFunc(valores ...int) float32{

	sum :=0
	for i:= 0; i< len(valores); i++{
		sum += valores[i]
	}
	return float32(sum/len(valores))
}

func maxFunc(valores ...int) float32{

	max:= valores[0]
	for i:= 1; i< len(valores); i++{
		if (valores[i] > max){
			max = valores[i]
		}
	}
	return float32(max)
}

func defaultFunc(grades ...int)(float32){
	return 0.0
}