package amazon

import (
	"github.com/amazinsellers/ecommerce-models/generic"
	"strings"
)

type OrderItem struct {
	ASIN                       string
	SellerSKU                  string
	OrderItemId                string
	Title                      string
	QuantityOrdered            int
	QuantityShipped            int
	ProductInfo                ProductInfoDetail
	PointsGranted              PointsGrantedDetail
	ItemPrice                  float64
	ShippingPrice              float64
	ItemTax                    float64
	ShippingTax                float64
	ShippingDiscount           float64
	ShippingDiscountTax        float64
	PromotionDiscount          float64
	PromotionDiscountTax       float64
	PromotionIds               PromotionIdList
	CODFee                     float64
	CODFeeDiscount             float64
	IsGift                     bool
	ConditionNote              string
	ConditionId                string
	ConditionSubtypeId         string
	ScheduledDeliveryStartDate string
	ScheduledDeliveryEndDate   string
	PriceDesignation           string
	TaxCollection              *TaxCollection
	SerialNumberRequired       bool
	IsTransparency             bool
	IossNumber                 string
	DeemedResellerCategory     string
}

type TaxCollection struct {
	Model            string
	ResponsibleParty string
}

type PromotionIdList []string

type PointsGrantedDetail struct {
	PointsNumber int
	PointsMonetaryValue float64
}
type ProductInfoDetail struct {
	NumberOfItems int
}

func (o *OrderItem) Generalise() *generic.OrderItem {
	return &generic.OrderItem{
		ChannelOrderItemId: o.OrderItemId,
		ChannelSKU:         o.ASIN,
		SellerSKU:          o.SellerSKU,
		QuantityOrdered:    float64(o.QuantityOrdered),
		QuantityShipped:    float64(o.QuantityShipped),

		ItemPrice:            o.ItemPrice,
		ItemTax:              o.ItemTax,
		PromotionDiscount:    o.PromotionDiscount,
		PromotionDiscountTax: o.PromotionDiscountTax,

		ShippingPrice:       o.ShippingPrice,
		ShippingTax:         o.ShippingTax,
		ShippingDiscount:    o.ShippingDiscount,
		ShippingDiscountTax: o.ShippingDiscountTax,

		PromotionIds: strings.Join(o.PromotionIds[:], ","),

		IsGift:                     o.IsGift,
		IsTaxWithheldByMarketplace: o.TaxCollection != nil,

		ScheduledDeliveryStartDate: o.ScheduledDeliveryStartDate,
		ScheduledDeliveryEndDate:   o.ScheduledDeliveryEndDate,

		IossNumber:             o.IossNumber,
		DeemedResellerCategory: o.DeemedResellerCategory,
	}
}
