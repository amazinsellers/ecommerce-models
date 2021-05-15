package generic

import (
	"time"
)

type Order struct {
	ChannelOrderId         string
	PurchaseDate           time.Time
	OrderStatus            string
	CurrencyCode		   string
	ItemsTotal             float64
	ShippingTotal          float64
	GrandTotal             float64
	IsBusinessOrder        bool
	Items                  []OrderItem
	SalesChannel           string
	OrderChannel           string
	ShipService            string
	FulfillmentChannel     string
	NumberOfItemsShipped   int
	NumberOfItemsUnshipped int
	OrderType              string
	EarliestShipDate       *time.Time
	LatestShipDate         *time.Time
	EarliestDeliveryDate   *time.Time
	LatestDeliveryDate     *time.Time
	ReplacedOrderId        string
	IsReplacementOrder     bool
	ShipToAddress          *Address
	PaymentMethod          string
	PaymentExecutionDetail []PaymentExecutionDetailItem
}

func (o *Order) AddItem(item OrderItem) {
	o.Items = append(o.Items, item)
	o.ItemsTotal += item.ItemPrice - item.PromotionDiscount
	o.ShippingTotal += item.ShippingPrice - item.ShippingDiscount
	o.GrandTotal = o.ItemsTotal + o.ShippingTotal
}

type PaymentExecutionDetailItem struct {
	Payment       float64
	PaymentMethod string
}
