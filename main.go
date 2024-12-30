package main

import (
	"fmt"
	"math"
)

func main() {
	const BMIPow = 2
	userHeight := 1.8
	userWeight := 75.0
	BMI := userWeight / math.Pow(userHeight, BMIPow)
	fmt.Print(BMI)
}
