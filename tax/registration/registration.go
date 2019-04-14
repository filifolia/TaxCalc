package registration

import (
	"TaxCalc/tax"
	"TaxCalc/tax/entertainment"
	"TaxCalc/tax/food"
	"TaxCalc/tax/tobacco"
)

func init() {
	//Initialize the registration of different type of tax calculators
	tax.Register(1, food.NewTaxCalculator())
	tax.Register(2, tobacco.NewTaxCalculator())
	tax.Register(3, entertainment.NewTaxCalculator())
}
