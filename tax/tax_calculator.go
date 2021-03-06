package tax

type TaxCalculator interface {
	Calculate(price float64) float64
	IsRefundable() string
	GetType() string
}
