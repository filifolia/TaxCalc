package models

//Item Model for the tax item object
type Item struct {
	Name       string  `orm:"pk" valid:"Required" json:"name"`
	TaxCode    int     `valid:"Required" json:"tax_code"`
	Price      float64 `json:"price"`
	Type       string  `orm:"-"`
	Refundable bool    `orm:"-"`
	Tax        float64 `orm:"-"`
	Amount     float64 `orm:"-"`
}
