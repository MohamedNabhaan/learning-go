package main 

import "math"
import "fmt"

func main(){
	const inflationRate = 2.5
	var investment = 100
	var interest = 5.6
	var years = 5

	var finalAmount = float64(investment) * math.Pow((1 + interest/100), float64(years))

	var finalAmountWithInflation = finalAmount * math.Pow((1 + inflationRate/100), float64(years))
	fmt.Println(finalAmountWithInflation)
	fmt.Println(finalAmount)
}