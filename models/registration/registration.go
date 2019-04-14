package registration

import (
	"TaxCalc/models"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	connString := "postgres://postgres:password@postgres/tax-db?connect_timeout=3&sslmode=disable"
	fmt.Println("Connecting to ", connString)

	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true

	orm.RegisterDataBase("default", "postgres", connString, 5, 5)

	orm.RegisterModel(new(models.Item))
}
