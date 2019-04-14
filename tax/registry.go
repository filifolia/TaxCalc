package tax

import (
	"errors"
	"strconv"
)

type Registry interface {
	GetTaxCalculators(code int) (TaxCalculator, error)
}

type registry struct{}

var taxcalculators = make(map[int]TaxCalculator)

func GetRegistry() Registry {
	return &registry{}
}

func Register(number int, taxcalculator TaxCalculator) {
	taxcalculators[number] = taxcalculator
}

func (*registry) GetTaxCalculators(code int) (TaxCalculator, error) {
	if factory, ok := taxcalculators[code]; ok {
		return factory, nil
	}
	return nil, errors.New("tax calculator not found with code " + strconv.Itoa(code))
}
