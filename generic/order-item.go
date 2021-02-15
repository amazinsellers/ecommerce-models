package generic

type OrderItem struct {
	ChannelOrderItemId string
	ChannelSKU         string
	SellerSKU          string
	QuantityOrdered    float64
	QuantityShipped    float64

	ItemPrice            float64
	ItemTax              float64
	PromotionDiscount    float64
	PromotionDiscountTax float64

	ShippingPrice       float64
	ShippingTax         float64
	ShippingDiscount    float64
	ShippingDiscountTax float64

	PromotionIds string

	IsGift                     bool
	IsTaxWithheldByMarketplace bool

	ScheduledDeliveryStartDate string
	ScheduledDeliveryEndDate   string

	IossNumber             string
	DeemedResellerCategory string
}
