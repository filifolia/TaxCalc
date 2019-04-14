package controllers

import (
	"TaxCalc/models"
	"TaxCalc/tax"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ItemController struct {
	beego.Controller
	itemOrmer models.ItemOrmer
	registry  tax.Registry
}

func (c *ItemController) Prepare() {
	c.registry = tax.GetRegistry()
	ormer := orm.NewOrm()
	c.itemOrmer = models.NewItemOrmer(ormer)
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type insertItemParams struct {
	Name    string  `json:"name"`
	TaxCode int     `json:"tax_code"`
	Price   float64 `json:"price"`
}

// @Param  params          body   {insertItemParams} true "params"
// @router /insert [post]
func (c *ItemController) Insert(params insertItemParams) *ErrorMessage {
	newItem := &models.Item{
		Name:    params.Name,
		Price:   params.Price,
		TaxCode: params.TaxCode,
	}

	err := c.itemOrmer.Create(newItem)
	if err != nil {
		errMsg := &ErrorMessage{Message: err.Error()}
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = errMsg
		c.ServeJSON()

		return errMsg
	}

	succMsg := &ErrorMessage{Message: ""}
	c.Data["json"] = succMsg
	c.ServeJSON()
	return succMsg
}

type ListMessage struct {
	Items         []*models.Item `json:"items"`
	PriceSubtotal float64        `json:"price_subtotal"`
	TaxSubtotal   float64        `json:"tax_subtotal"`
	GrandTotal    float64        `json:"grand_total"`
}

// @router /get [get]
func (c *ItemController) GetList() (*ListMessage, *ErrorMessage) {
	items, err := c.itemOrmer.GetItems()
	if err != nil {
		errMsg := &ErrorMessage{Message: err.Error()}
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = errMsg
		c.ServeJSON()

		return nil, errMsg
	}

	//Fill up the other tax item fields
	var totalPrice float64
	var totalTax float64
	for _, item := range items {
		taxCalc, err := c.registry.GetTaxCalculators(item.TaxCode)
		if err != nil {
			log.Println("Error tax calc: ", err)
			continue
		}

		item.Tax = taxCalc.Calculate(item.Price)
		item.Refundable = taxCalc.IsRefundable()
		item.Type = taxCalc.GetType()
		item.Amount = item.Tax + item.Price
		totalPrice += item.Price
		totalTax += item.Tax
	}

	var resp = &ListMessage{
		Items:         items,
		PriceSubtotal: totalPrice,
		TaxSubtotal:   totalTax,
		GrandTotal:    totalPrice + totalTax,
	}

	c.Data["json"] = resp
	c.ServeJSON()

	return resp, nil
}
