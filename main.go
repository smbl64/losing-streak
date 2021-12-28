package main

import (
	"fmt"
	"math"
)

func main() {
	totalTrades := 50
	minNumLoses := 2
	maxNumLoses := 11
	minWinRate := 5
	maxWinRate := 95

	fmt.Printf("Probability of seeing at least (X) consecutive losing trades within a %d-trade period\n", totalTrades)
	fmt.Printf("\n")

	// To align with the win rates below
	fmt.Print("          ")
	for numLoses := minNumLoses; numLoses <= maxNumLoses; numLoses++ {
		fmt.Printf("%5d   ", numLoses)
	}
	fmt.Printf("\n")

	for winRate := minWinRate; winRate <= maxWinRate; winRate += 5 {
		fmt.Printf("%3d%%      ", winRate)

		for numLoses := minNumLoses; numLoses <= maxNumLoses; numLoses++ {
			result := 100 * calculate(numLoses, totalTrades, float64(winRate)/100)
			fmt.Printf("%5.1f   ", result)
		}

		fmt.Printf("\n")
	}
}

func calculate(numLoses, totalTrades int, winRate float64) float64 {
	// Formula taken from:
	// https://www.forexfactory.com/thread/post/4336468#post4336468
	// It matches the table shared on the Internet and FTMO
	return 1 - math.Pow(1-math.Pow(1-winRate, float64(numLoses)), float64(totalTrades-(numLoses-1)))
}
