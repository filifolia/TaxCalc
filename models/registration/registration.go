package registration

import (
	"TaxCalc/models"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	envVar := "dev"
	osEnv := os.Getenv("TAX_ENV")
	if osEnv != "" {
		envVar = osEnv
	}

	log.Println("Running in environment:" + envVar)
	orm.RegisterDriver("postgres", orm.DRPostgres)

	dbName := beego.AppConfig.String("databaseName")
	if dbName == "" {
		log.Println("Fail to get DB name config, default to test: " + envVar)
		dbName = "tax-db-test"
	}

	connString := fmt.Sprintf("postgres://postgres:password@postgres/%s?connect_timeout=3&sslmode=disable", dbName)
	fmt.Println("Connecting to ", connString)

	orm.DefaultTimeLoc = time.UTC
	orm.Debug = true

	orm.RegisterDataBase("default", "postgres", connString, 5, 5)

	orm.RegisterModel(new(models.Item))
}
