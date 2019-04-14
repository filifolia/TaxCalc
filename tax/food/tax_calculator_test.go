package food_test

import (
	"TaxCalc/tax/food"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTaxCalculator(t *testing.T) {
	taxCalc := food.NewTaxCalculator()
	Convey("Calculate", t, func() {
		result := taxCalc.Calculate(500)
		So(result, ShouldEqual, 50)

		result = taxCalc.Calculate(1)
		So(result, ShouldEqual, 0.1)
	})

	Convey("IsRefundable", t, func() {
		result := taxCalc.IsRefundable()
		So(result, ShouldEqual, "Yes")
	})
}
