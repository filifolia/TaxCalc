package routers

import (
	"TaxCalc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/item",
			beego.NSInclude(
				&controllers.ItemController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
