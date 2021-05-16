package amazon

import "github.com/amazinsellers/ecommerce-models/generic"

type Address struct {
	Name string

	AddressLine1 string
	AddressLine2 string
	AddressLine3 string

	City          string
	County        string
	District      string
	StateOrRegion string
	Municipality  string

	PostalCode  string
	CountryCode string

	Phone       string
	AddressType AddressType
}

type AddressType int

const (
	Residential AddressType = iota
	Commercial
)

func (o *Address) Generalise() *generic.Address {
	return &generic.Address{
		Name: o.Name,

		AddressLine1: o.AddressLine1,
		AddressLine2: o.AddressLine2,
		AddressLine3: o.AddressLine3,

		City: o.City,
		County: o.County,
		District: o.District,
		StateOrRegion: o.StateOrRegion,

		Municipality: o.Municipality,
		PostalCode: o.PostalCode,
		CountryCode: o.CountryCode,

		Phone: o.Phone,
		AddressType: generic.AddressType(int(o.AddressType)),
	}
}
