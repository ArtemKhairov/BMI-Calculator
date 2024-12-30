package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPow = 2
	var userHeight float64
	var userWeight float64
	fmt.Println("Body mass index calculator")
	fmt.Print("Введите рост:")
	fmt.Scan(&userHeight)
	fmt.Print("Введите вес:")
	fmt.Scan(&userWeight)
	BMI := userWeight / math.Pow(userHeight, BMIPow)
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", BMI)
	fmt.Print(result)
}
