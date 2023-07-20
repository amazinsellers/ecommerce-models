package amazon

import (
	"encoding/json"
	"fmt"
	"github.com/amazinsellers/ecommerce-models/amazon"
	"io/ioutil"
	"testing"
)

func TestProduct(t *testing.T) {
	file, err := ioutil.ReadFile("./amazon-product.json")
	if err != nil {
		t.Error("amazon-product.json file not readable")
	}

	jsonContent := string(file)

	productArr := &amazon.ProductArray{}
	err = productArr.Unmarshal(jsonContent)

	if err != nil {
		t.Error(err)
		return
	}

	marshal, _ := json.Marshal(productArr)
	fmt.Println(string(marshal))
}
