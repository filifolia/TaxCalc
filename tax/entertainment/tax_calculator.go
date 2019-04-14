package entertainment

import "TaxCalc/tax"

type taxCalculator struct {
}

func NewTaxCalculator() tax.TaxCalculator {
	return &taxCalculator{}
}

func (t *taxCalculator) Calculate(price float64) float64 {
	if price >= float64(100) {
		return (price - float64(100)) / float64(100)
	}

	return 0
}

func (t *taxCalculator) IsRefundable() string {
	return "No"
}

func (t *taxCalculator) GetType() string {
	return "Entertainment"
}
