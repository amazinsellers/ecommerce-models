package generic

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
