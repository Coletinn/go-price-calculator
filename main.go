package main

import (
	"fmt"
	"price-calculator/prices"
	"price-calculator/cmd"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		cmd := cmd.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmd, taxRate)
		priceJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
	fmt.Println()
}