package amazon

import (
	"encoding/json"
	"github.com/amazinsellers/ecommerce-models/generic"
	"time"
)

type Product struct {
	ItemName                string    `json:"item-name"`
	ItemDescription         string    `json:"item-description"`
	ListingId               string    `json:"listing-id"`
	SellerSKU               string    `json:"seller-sku"`
	Price                   float64   `json:"price"`
	Quantity                int64     `json:"quantity"`
	OpenDate                time.Time `json:"open-date"`
	ImageUrl                string    `json:"image-url"`
	ItemIsMarketplace       string    `json:"item-is-marketplace"`
	ProductIdType           int64     `json:"product-id-type"`
	ZshopShippingFee        string    `json:"zshop-shipping-fee"`
	ItemNote                string    `json:"item-note"`
	ItemCondition           int64     `json:"item-condition"`
	ZshopCategory1          string    `json:"zshop-category-1"`
	ZshopBrowsePath         string    `json:"zshop-browse-path"`
	ZshopStorefrontFeature  string    `json:"zshop-storefront-feature"`
	Asin1                   string    `json:"asin1"`
	Asin2                   string    `json:"asin2"`
	Asin3                   string    `json:"asin3"`
	WillShipInternationally string    `json:"will-ship-internationally"`
	ExpeditedShipping       string    `json:"expedited-shipping"`
	ZshopBoldface           string    `json:"zshop-boldface"`
	ProductId               string    `json:"product-id"`
	BidForFeaturedPlacement string    `json:"bid-for-featured-placement"`
	AddDelete               string    `json:"add-delete"`
	PendingQuantity         int64     `json:"pending-quantity"`
	FulfillmentChannel      string    `json:"fulfillment-channel"`
	MerchantShippingGroup   string    `json:"merchant-shipping-group"`
	Status                  string    `json:"status"`
}

func (o *Product) Generalise(currencyCode string) *generic.Product {
	return &generic.Product{
		SKU:         o.SellerSKU,
		ChannelSKU:  o.Asin1,
		Title:       o.ItemName,
		Description: o.ItemDescription,
		ImageUrls:   o.ImageUrl,
		Prices: []generic.Price{
			{
				Amount:   o.Price,
				Currency: currencyCode,
			},
		},
	}
}

func (o *ProductArray) Unmarshal(productsStr string) error {
	err := json.Unmarshal([]byte(productsStr), o)

	if err != nil {
		return err
	}

	return nil
}

type ProductArray []Product

func (o ProductArray) Generalise(currencyCode string) []generic.Product {
	genericProducts := make([]generic.Product, 0)

	for _, aProduct := range o {
		genericProducts = append(genericProducts, generic.Product{
			SKU:         aProduct.SellerSKU,
			ChannelSKU:  aProduct.Asin1,
			Title:       aProduct.ItemName,
			Description: aProduct.ItemDescription,
			ImageUrls:   aProduct.ImageUrl,
			Prices: []generic.Price{
				{
					Amount:   aProduct.Price,
					Currency: currencyCode,
				},
			},
		})
	}

	return genericProducts
}
