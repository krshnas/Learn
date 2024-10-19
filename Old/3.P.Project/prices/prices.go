package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxtIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_price"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManger          iomanager.IOManager `json:"-"` // here - in json struct tag telling json to ignore this in output
}

func NewTaxIncludedPriceJob(ioManager iomanager.IOManager, taxRate float64) *TaxtIncludedPriceJob {
	return &TaxtIncludedPriceJob{
		IOManger:    ioManager,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *TaxtIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()
	if err != nil {
		// return err
		errorChan <- err
		return
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManger.WriteResult(job)
	doneChan <- true
}

func (job *TaxtIncludedPriceJob) LoadData() error {
	lines, err := job.IOManger.ReadLines()
	if err != nil {
		return err
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil
}
