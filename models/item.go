package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

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
		return errors.New(valid.Errors[0].Error())
	}
	_, err = i.ormer.Insert(item)
	return err
}

func (i *itemOrmer) GetItems() ([]*Item, error) {
	var items []*Item
	_, err := i.ormer.QueryTable("item").Limit(-1).All(&items)
	return items, err
}
