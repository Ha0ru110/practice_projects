package main

import (
	"price-calculator.com/haorui/cmdmanager"
	"price-calculator.com/haorui/prices"
)

func main() {
	taxRates := []float64{1, 2.5, 3}
	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			return
		}
	}
}
