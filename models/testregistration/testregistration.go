package testregistration

import (
	"fmt"
	"path/filepath"

	"github.com/astaxie/beego"
)

func init() {
	beego.TestBeegoInit(getAppBasePath())
}

func getAppBasePath() string {
	basePath, _ := filepath.Abs(".")
	for filepath.Base(basePath) != "TaxCalc" {
		basePath = filepath.Dir(basePath)
	}
	fmt.Println(basePath)
	return basePath
}
