package e2e_test

import (
	_ "TaxCalc/models/registration"
	"log"

	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/astaxie/beego/orm"
	. "github.com/smartystreets/goconvey/convey"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

//END TO END TEST by invoking http calls to the API itself
func TestEndToEnd(t *testing.T) {
	Convey("The API should be able to create two items and return them", t, func() {
		ormer := orm.NewOrm()
		//Reset the test DB first
		_, err := ormer.Raw("TRUNCATE TABLE item CASCADE;").Exec()
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
		var bodyReader io.Reader
		var msg ErrorMessage
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		//////////////////First Item
		body := `{
			"tax_code" : 2,
			"name" : "TEST",
			"price" : 100.5
		}`
		bodyReader = bytes.NewReader([]byte(body))

		req, err := http.NewRequest("POST", "http://api:8080/v1/item/insert", bodyReader)
		So(err, ShouldBeNil)

		resp, err := client.Do(req)
		So(err, ShouldBeNil)

		respBytes, err := ioutil.ReadAll(resp.Body)
		So(err, ShouldBeNil)

		log.Printf("RESPONSE: %s", string(respBytes))
		resp.Body.Close()

		err = json.Unmarshal(respBytes, &msg)
		So(err, ShouldBeNil)
		So(msg.Message, ShouldEqual, "")

		/////////////////Second Item
		body = `{
			"tax_code" : 1,
			"name" : "TEST2",
			"price" : 150
		}`
		bodyReader = bytes.NewReader([]byte(body))

		req, err = http.NewRequest("POST", "http://api:8080/v1/item/insert", bodyReader)
		So(err, ShouldBeNil)

		resp, err = client.Do(req)
		So(err, ShouldBeNil)

		respBytes, err = ioutil.ReadAll(resp.Body)
		So(err, ShouldBeNil)

		log.Printf("RESPONSE: %s", string(respBytes))
		resp.Body.Close()

		err = json.Unmarshal(respBytes, &msg)
		So(err, ShouldBeNil)
		So(msg.Message, ShouldEqual, "")

	})

}
