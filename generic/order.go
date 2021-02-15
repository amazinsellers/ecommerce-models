package generic

import (
	"time"
)

type Order struct {
	ChannelOrderId         string
	PurchaseDate           time.Time
	OrderStatus            string
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
	EarliestShipDate       time.Time
	LatestShipDate         time.Time
	EarliestDeliveryDate   time.Time
	LatestDeliveryDate     time.Time
	ReplacedOrderId        string
	IsReplacementOrder     bool

	PaymentMethod          string
	PaymentExecutionDetail []PaymentExecutionDetailItem
}

type PaymentExecutionDetailItem struct {
	Payment       float64
	PaymentMethod string
}
