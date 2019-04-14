package food

import "TaxCalc/tax"

type taxCalculator struct {
}

func NewTaxCalculator() tax.TaxCalculator {
	return &taxCalculator{}
}

func (t *taxCalculator) Calculate(price float64) float64 {
	return (price / 10)
}

func (t *taxCalculator) IsRefundable() string {
	return "Yes"
}
