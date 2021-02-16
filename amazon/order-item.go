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
	ItemPrice                  Money
	ShippingPrice              Money
	ItemTax                    Money
	ShippingTax                Money
	ShippingDiscount           Money
	ShippingDiscountTax        Money
	PromotionDiscount          Money
	PromotionDiscountTax       Money
	PromotionIds               PromotionIdList
	CODFee                     Money
	CODFeeDiscount             Money
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

		ItemPrice:            o.ItemPrice.Amount,
		ItemTax:              o.ItemTax.Amount,
		PromotionDiscount:    o.PromotionDiscount.Amount,
		PromotionDiscountTax: o.PromotionDiscountTax.Amount,

		ShippingPrice:       o.ShippingPrice.Amount,
		ShippingTax:         o.ShippingTax.Amount,
		ShippingDiscount:    o.ShippingDiscount.Amount,
		ShippingDiscountTax: o.ShippingDiscountTax.Amount,

		PromotionIds: strings.Join(o.PromotionIds[:], ","),

		IsGift:                     o.IsGift,
		IsTaxWithheldByMarketplace: o.TaxCollection != nil,

		ScheduledDeliveryStartDate: o.ScheduledDeliveryStartDate,
		ScheduledDeliveryEndDate:   o.ScheduledDeliveryEndDate,

		IossNumber:             o.IossNumber,
		DeemedResellerCategory: o.DeemedResellerCategory,
	}
}
