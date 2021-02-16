package amazon

import (
	"github.com/amazinsellers/ecommerce-models/generic"
	"time"
)

type Order struct {
	AmazonOrderId                   string
	SellerOrderId                   string
	PurchaseDate                    string
	LastUpdateDate                  string
	OrderStatus                     string
	FulfillmentChannel              string
	SalesChannel                    string
	OrderChannel                    string
	ShipServiceLevel                string
	OrderTotal                      Money
	NumberOfItemsShipped            int
	NumberOfItemsUnshipped          int
	PaymentExecutionDetail          PaymentExecutionDetailItemList
	PaymentMethod                   string
	PaymentMethodDetails            PaymentMethodDetailItemList
	MarketplaceId                   string
	ShipmentServiceLevelCategory    string
	EasyShipShipmentStatus          string
	CbaDisplayableShippingLabel     string
	OrderType                       string
	EarliestShipDate                string
	LatestShipDate                  string
	EarliestDeliveryDate            string
	LatestDeliveryDate              string
	IsBusinessOrder                 bool
	IsPrime                         bool
	IsPremiumOrder                  bool
	IsGlobalExpressEnabled          bool
	ReplacedOrderId                 string
	IsReplacementOrder              bool
	PromiseResponseDueDate          string
	IsEstimatedShipDateSet          bool
	IsSoldByAB                      bool
	AssignedShipFromLocationAddress Address
	FulfillmentInstruction          FulfillmentInstruction

	Items							[]OrderItem
}

type FulfillmentInstruction struct {
	FulfillmentSupplySourceId string
}

type PaymentMethodDetailItemList []string

type PaymentExecutionDetailItemList []generic.PaymentExecutionDetailItem

type Address struct {
	Name          string
	AddressLine1  string
	AddressLine2  string
	AddressLine3  string
	City          string
	County        string
	District      string
	StateOrRegion string
	Municipality  string
	PostalCode    string
	CountryCode   string
	Phone         string
	AddressType   string
}

func (o *Order) Generalise() *generic.Order {
	layout := "2006-01-02T15:04:05.000Z"
	purchaseDate, err := time.Parse(layout, o.PurchaseDate)
	if err != nil {
		purchaseDate = time.Now()
	}

	earliestShipDate, err := time.Parse(layout, o.EarliestShipDate)
	if err != nil {
		earliestShipDate = time.Now()
	}

	latestShipDate, err := time.Parse(layout, o.LatestShipDate)
	if err != nil {
		latestShipDate = time.Now()
	}

	latestDeliveryDate, err := time.Parse(layout, o.LatestDeliveryDate)
	if err != nil {
		latestDeliveryDate = time.Now()
	}

	earliestDeliveryDate, err := time.Parse(layout, o.EarliestDeliveryDate)
	if err != nil {
		earliestDeliveryDate = time.Now()
	}

	theOrder := &generic.Order{
		ChannelOrderId:         o.AmazonOrderId,
		PurchaseDate:           purchaseDate,
		OrderStatus:            o.OrderStatus,
		IsBusinessOrder:        o.IsBusinessOrder,
		SalesChannel:           o.SalesChannel,
		OrderChannel:           o.OrderChannel,
		ShipService:            o.ShipServiceLevel + " " + o.ShipmentServiceLevelCategory,
		FulfillmentChannel:     o.FulfillmentChannel,
		NumberOfItemsShipped:   o.NumberOfItemsShipped,
		NumberOfItemsUnshipped: o.NumberOfItemsUnshipped,
		OrderType:              o.OrderType,
		EarliestShipDate:       earliestShipDate,
		LatestShipDate:         latestShipDate,
		LatestDeliveryDate:     latestDeliveryDate,
		EarliestDeliveryDate:   earliestDeliveryDate,
		ReplacedOrderId:        o.ReplacedOrderId,
		IsReplacementOrder:     o.IsReplacementOrder,
		PaymentMethod:          o.PaymentMethod,
		PaymentExecutionDetail: o.PaymentExecutionDetail,
		CurrencyCode:           o.OrderTotal.CurrencyCode,
	}

	for _, v := range o.Items {
		theOrder.AddItem(*v.Generalise())
	}

	return theOrder
}
