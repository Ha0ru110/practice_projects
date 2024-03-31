package prices

import (
	"errors"
	"fmt"
	"price-calculator.com/haorui/section-9-price-calc/conversion"
	"price-calculator.com/haorui/section-9-price-calc/iomanager"
)

type TaxIncludedPricesJob struct {
	IoManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices: []float64{1, 3, 3, 4},
		TaxRate:     taxRate,
		IoManager:   iom,
	}
}

func (job *TaxIncludedPricesJob) Process() error {
	err := job.readPrices()
	if err != nil {
		return err
	}
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	err = job.IoManager.WriteResult(job)
	if err != nil {
		return errors.New("failed to write results")
	}
	return nil
}

func (job *TaxIncludedPricesJob) readPrices() error {
	lines, err := job.IoManager.ReadLines()
	if err != nil {
		//if mysterious error, check here
		fmt.Println("reading data failed")
		fmt.Println(err)
		if err != nil {
			fmt.Println("panic")
		}
		return errors.New("failure at 49")

	}
	prices, err := conversion.Str2Float(lines)
	if err != nil {
		//if mysterious error, check here
		fmt.Println("converting data failed")
		fmt.Println(err)
		if err != nil {
			fmt.Println("panic")
		}
		return errors.New("failure at 60")

	}
	job.InputPrices = prices
	return nil
}
