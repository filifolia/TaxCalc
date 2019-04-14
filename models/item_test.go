package models_test

import (
	. "TaxCalc/models"
	_ "TaxCalc/models/registration"
	"testing"

	"github.com/astaxie/beego/orm"

	. "github.com/smartystreets/goconvey/convey"
)

func TestItemOrmer(t *testing.T) {
	ormer := orm.NewOrm()
	//Reset the test DB first
	_, err := ormer.Raw("TRUNCATE TABLE item CASCADE;").Exec()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	taxCalc := NewItemOrmer(ormer)
	Convey("CRUD", t, func() {
		item := &Item{
			Name:    "itemTEST",
			Price:   100,
			TaxCode: 1,
		}

		err := taxCalc.Create(item)
		So(err, ShouldBeNil)

		item.Price = 200
		err = taxCalc.Create(item)
		So(err, ShouldNotBeNil)

		item2 := &Item{
			Name:    "itemTEST2",
			Price:   500,
			TaxCode: 2,
		}
		err = taxCalc.Create(item2)
		So(err, ShouldBeNil)

		//Test to get all 2 items
		items, err := taxCalc.GetItems()
		So(err, ShouldBeNil)
		So(items, ShouldHaveLength, 2)
		So(items[0].Name, ShouldEqual, "itemTEST")
		So(items[1].Name, ShouldEqual, "itemTEST2")

	})

}
