package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Body mass index calculator")
	for {
		const BMIPow = 2
		userHeight, userWeight := getUserParams()
		BMI := calculateBMI(userHeight, userWeight, BMIPow)
		bmiClassification := classifyBMI(BMI)
		fmt.Println(bmiClassification)
		result := fmt.Sprintf("Your body mass index: %.0f", BMI)
		fmt.Println(result)
		isRepeate := checkRepeateCalculation()
		if !isRepeate {
			break
		}
	}
}

func classifyBMI(bmi float64) string {
	switch {
	case bmi < 16:
		return "Severe underweight"
	case bmi >= 16 && bmi <= 16.9:
		return "Underweight"
	case bmi >= 17 && bmi <= 18.4:
		return "Slightly underweight"
	case bmi >= 18.5 && bmi <= 24.9:
		return "Normal weight"
	case bmi >= 25 && bmi <= 29.9:
		return "Overweight"
	case bmi >= 30 && bmi <= 34.9:
		return "Obesity Class 1"
	case bmi >= 35 && bmi <= 39.9:
		return "Obesity Class 2"
	default: // bmi >= 40
		return "Obesity Class 3"
	}
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

func checkRepeateCalculation() bool {
	var userChoice string
	fmt.Println("Do you want to repeate? y/n:")
	fmt.Scan(&userChoice)
	if userChoice == "Y" || userChoice == "y" {
		return true
	}
	return false
}
