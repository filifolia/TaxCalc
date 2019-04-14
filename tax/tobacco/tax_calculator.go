package tobacco

import "TaxCalc/tax"

type taxCalculator struct {
}

func NewTaxCalculator() tax.TaxCalculator {
	return &taxCalculator{}
}

func (t *taxCalculator) Calculate(price float64) float64 {
	return float64(10) + (price / 50)
}

func (t *taxCalculator) IsRefundable() string {
	return "No"
}

func (t *taxCalculator) GetType() string {
	return "Tobacco"
}
