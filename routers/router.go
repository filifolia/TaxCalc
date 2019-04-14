package routers

import (
	"TaxCalc/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/item", &controllers.ItemController{})
}
