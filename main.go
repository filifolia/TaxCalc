package main

import (
	_ "TaxCalc/models/registration"
	_ "TaxCalc/routers"
	_ "TaxCalc/tax/registration"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
