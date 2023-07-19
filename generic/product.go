package generic

import (
	"encoding/json"
	"log"
	"math"
)

type Product struct {
	SKU         string
	ChannelSKU  string
	Title       string
	Description string
	CustomerId  string
	ImageUrls   string

	Prices []Price
}

type ProductList struct {
	TotalItems   int64
	CurrentPage  int
	TotalPages   int64
	ItemsPerPage int
	Products     *[]Product
}

func (o *ProductList) ToJson() string {
	if o.ItemsPerPage == 0 {
		o.ItemsPerPage = 50
	}

	log.Println(o.TotalItems,
		int64(o.ItemsPerPage), float64(o.TotalItems)/float64(o.ItemsPerPage),
		float64(o.TotalItems)/float64(o.ItemsPerPage),
		math.Ceil(float64(o.TotalItems)/float64(o.ItemsPerPage)))

	o.TotalPages = int64(math.Ceil(float64(o.TotalItems) / float64(o.ItemsPerPage)))
	bytes, _ := json.Marshal(o)
	return string(bytes)
}
