package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

//Item Model for the tax item object
type Item struct {
	Name       string  `orm:"column(name);pk" valid:"Required" json:"name"`
	TaxCode    int     `orm:"column(tax_code)" valid:"Required" json:"tax_code"`
	Price      float64 `orm:"column(price)" json:"price"`
	Type       string  `orm:"-" json:"type"`
	Refundable string  `orm:"-" json:"refundable"`
	Tax        float64 `orm:"-" json:"tax"`
	Amount     float64 `orm:"-" json:"amount"`
}

type ItemOrmer interface {
	GetItems() ([]*Item, error)
	Create(item *Item) error
}

type itemOrmer struct {
	ormer orm.Ormer
}

func NewItemOrmer(ormer orm.Ormer) ItemOrmer {
	return &itemOrmer{ormer: ormer}
}

func (i *itemOrmer) Create(item *Item) error {
	valid := validation.Validation{}
	ok, err := valid.Valid(item)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New(valid.Errors[0].Error() + valid.Errors[0].Field)
	}
	_, err = i.ormer.Raw("INSERT INTO item(name, tax_code, price) VALUES(?, ?, ?);", item.Name, item.TaxCode, item.Price).Exec()
	if err != nil {
		return err
	}

	return nil
}

func (i *itemOrmer) GetItems() ([]*Item, error) {
	var items []*Item
	_, err := i.ormer.QueryTable("item").Limit(-1).All(&items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
