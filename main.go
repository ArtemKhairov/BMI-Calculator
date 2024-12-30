package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPow = 2
	fmt.Println("Body mass index calculator")
	userHeight, userWeight := getUserParams()
	BMI := calculateBMI(userHeight, userWeight, BMIPow)
	if BMI < 16 {
		fmt.Print("Your weigh is to low")
	}
	result := fmt.Sprintf("Your body mass index: %.0f", BMI)
	fmt.Print(result)
}

func calculateBMI(userHeight, userWeight, BMIPow float64) float64 {
	return userWeight / math.Pow(userHeight, BMIPow)
}

func getUserParams() (float64, float64) {
	var userHeight float64
	var userWeight float64
	fmt.Print("Enter your height:")
	fmt.Scan(&userHeight)
	fmt.Print("Enter your weight:")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}
