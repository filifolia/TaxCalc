package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["TaxCalc/controllers:ItemController"] = append(beego.GlobalControllerRouter["TaxCalc/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetList",
            Router: `/get`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["TaxCalc/controllers:ItemController"] = append(beego.GlobalControllerRouter["TaxCalc/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Insert",
            Router: `/insert`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("params", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

}
