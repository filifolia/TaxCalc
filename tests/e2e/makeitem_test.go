package e2e_test

import (
	"TaxCalc/models"
	_ "TaxCalc/models/registration"
	"log"

	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type ListMessage struct {
	Items         []*models.Item `json:"items"`
	PriceSubtotal float64        `json:"price_subtotal"`
	TaxSubtotal   float64        `json:"tax_subtotal"`
	GrandTotal    float64        `json:"grand_total"`
}

//END TO END TEST by invoking http calls to the API itself
func TestEndToEnd(t *testing.T) {
	Convey("The API should be able to create two items and return them", t, func() {
		var bodyReader io.Reader
		var msg ErrorMessage
		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		//////////////////First Item
		body := `{
			"tax_code" : 2,
			"name" : "TEST1",
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

		///////////////Second Item
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

		///////////////Get Item List
		var listMsg ListMessage
		req, err = http.NewRequest("GET", "http://api:8080/v1/item/get", nil)
		So(err, ShouldBeNil)

		resp, err = client.Do(req)
		So(err, ShouldBeNil)

		respBytes, err = ioutil.ReadAll(resp.Body)
		So(err, ShouldBeNil)

		log.Printf("RESPONSE: %s", string(respBytes))
		resp.Body.Close()

		err = json.Unmarshal(respBytes, &listMsg)
		So(err, ShouldBeNil)

	})

}
