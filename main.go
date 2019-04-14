package main

import (
	_ "TaxCalc/models/registration"
	_ "TaxCalc/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
