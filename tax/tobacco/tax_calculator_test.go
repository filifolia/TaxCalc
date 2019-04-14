package tobacco_test

import (
	"TaxCalc/tax/tobacco"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTaxCalculator(t *testing.T) {
	taxCalc := tobacco.NewTaxCalculator()
	Convey("Calculate", t, func() {
		result := taxCalc.Calculate(100)
		So(result, ShouldEqual, 12)

		result = taxCalc.Calculate(1)
		So(result, ShouldEqual, 10.02)
	})

	Convey("IsRefundable", t, func() {
		result := taxCalc.IsRefundable()
		So(result, ShouldEqual, "No")
	})

	Convey("GetType", t, func() {
		result := taxCalc.GetType()
		So(result, ShouldEqual, "Tobacco")
	})
}
