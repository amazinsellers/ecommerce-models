package amazon

import (
	"encoding/json"
	"github.com/amazinsellers/ecommerce-models/generic"
	"strconv"
)

type Product struct {
	ItemName                string `json:"item-name"`
	ItemDescription         string `json:"item-description"`
	ListingId               string `json:"listing-id"`
	SellerSKU               string `json:"seller-sku"`
	OpenDate                string `json:"open-date"`
	ImageUrl                string `json:"image-url"`
	ItemIsMarketplace       string `json:"item-is-marketplace"`
	ZshopShippingFee        string `json:"zshop-shipping-fee"`
	ItemNote                string `json:"item-note"`
	ZshopCategory1          string `json:"zshop-category-1"`
	ZshopBrowsePath         string `json:"zshop-browse-path"`
	ZshopStorefrontFeature  string `json:"zshop-storefront-feature"`
	Asin1                   string `json:"asin1"`
	Asin2                   string `json:"asin2"`
	Asin3                   string `json:"asin3"`
	WillShipInternationally string `json:"will-ship-internationally"`
	ExpeditedShipping       string `json:"expedited-shipping"`
	ZshopBoldface           string `json:"zshop-boldface"`
	ProductId               string `json:"product-id"`
	BidForFeaturedPlacement string `json:"bid-for-featured-placement"`
	AddDelete               string `json:"add-delete"`
	FulfillmentChannel      string `json:"fulfillment-channel"`
	MerchantShippingGroup   string `json:"merchant-shipping-group"`
	Status                  string `json:"status"`

	Price           string `json:"price"`
	Quantity        string `json:"quantity"`
	ProductIdType   string `json:"product-id-type"`
	ItemCondition   string `json:"item-condition"`
	PendingQuantity string `json:"pending-quantity"`
}

func (o *Product) Generalise(currencyCode string) (*generic.Product, error) {
	price, err := strconv.ParseFloat(o.Price, 64)

	if err != nil {
		return nil, err
	}

	return &generic.Product{
		SKU:         o.SellerSKU,
		ChannelSKU:  o.Asin1,
		Title:       o.ItemName,
		Description: o.ItemDescription,
		ImageUrls:   o.ImageUrl,
		Prices: []generic.Price{
			{
				Amount:   price,
				Currency: currencyCode,
			},
		},
	}, nil
}

func (o *ProductArray) Unmarshal(productsStr string) error {
	err := json.Unmarshal([]byte(productsStr), o)

	if err != nil {
		return err
	}

	return nil
}

type ProductArray []Product

func (o ProductArray) Generalise(currencyCode string) ([]generic.Product, error) {
	genericProducts := make([]generic.Product, 0)

	for _, aProduct := range o {
		aGenericProduct, err := aProduct.Generalise(currencyCode)
		if err != nil {
			return nil, err
		}

		genericProducts = append(genericProducts, *aGenericProduct)
	}

	return genericProducts, nil
}
