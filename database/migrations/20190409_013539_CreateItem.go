package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateItem_20190409_013539 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateItem_20190409_013539{}
	m.Created = "20190409_013539"

	migration.Register("CreateItem_20190409_013539", m)
}

// Run the migrations
func (m *CreateItem_20190409_013539) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE item (
		name varchar(255) PRIMARY KEY,
		tax_code integer,
		price float
		);`)
}

// Reverse the migrations
func (m *CreateItem_20190409_013539) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`DROP TABLE item;`)
}
