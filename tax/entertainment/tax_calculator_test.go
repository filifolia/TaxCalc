package entertainment_test

import (
	"TaxCalc/tax/entertainment"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTaxCalculator(t *testing.T) {
	taxCalc := entertainment.NewTaxCalculator()
	Convey("Calculate", t, func() {
		Convey("When it is taxable", func() {
			result := taxCalc.Calculate(500)
			So(result, ShouldEqual, 4)

			result = taxCalc.Calculate(101)
			So(result, ShouldEqual, 0.01)
		})

		Convey("When it is not appliable", func() {
			result := taxCalc.Calculate(99)
			So(result, ShouldEqual, 0)

			result = taxCalc.Calculate(100)
			So(result, ShouldEqual, 0)
		})

	})

	Convey("IsRefundable", t, func() {
		result := taxCalc.IsRefundable()
		So(result, ShouldEqual, "No")
	})
}
