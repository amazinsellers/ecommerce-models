package amazon

type Channel string

type ChannelProp struct {
	CurrencyCode  string
	Timezone      string
	MarketplaceId string
}

var ChannelMap = map[string]ChannelProp{
	"amazon-ae": {
		CurrencyCode:  "AED",
		Timezone:      "Asia/Dubai",
		MarketplaceId: "A2VIGQ35RCS4UG",
	},
	"amazon-de": {
		CurrencyCode:  "EUR",
		Timezone:      "Europe/Berlin",
		MarketplaceId: "A1PA6795UKMFR9",
	},
	"amazon-eg": {
		CurrencyCode:  "EGP",
		Timezone:      "Africa/Cairo",
		MarketplaceId: "ARBP9OOSHTCHU",
	},
	"amazon-es": {
		CurrencyCode:  "EUR",
		Timezone:      "Europe/Madrid",
		MarketplaceId: "A1RKKUPIHCS9HS",
	},
	"amazon-fr": {
		CurrencyCode:  "EUR",
		Timezone:      "Europe/Paris",
		MarketplaceId: "A13V1IB3VIYZZH",
	},
	"amazon-gb": {
		CurrencyCode:  "GBP",
		Timezone:      "Europe/London",
		MarketplaceId: "A1F83G8C2ARO7P",
	},
	"amazon-in": {
		CurrencyCode:  "INR",
		Timezone:      "Asia/Kolkata",
		MarketplaceId: "A21TJRUUN4KGV",
	},
	"amazon-nl": {
		CurrencyCode:  "EUR",
		Timezone:      "Europe/Amsterdam",
		MarketplaceId: "A1805IZSGTT6HS",
	},
	"amazon-pl": {
		CurrencyCode:  "PLN",
		Timezone:      "Europe/Warsaw",
		MarketplaceId: "A1C3SOZRARQ6R3",
	},
	"amazon-sa": {
		CurrencyCode:  "EUR",
		Timezone:      "Asia/Riyadh",
		MarketplaceId: "A17E79C6D8DWNP",
	},
	"amazon-se": {
		CurrencyCode:  "EUR",
		Timezone:      "Europe/Stockholm",
		MarketplaceId: "A2NODRKZP88ZB9",
	},
	"amazon-tr": {
		CurrencyCode:  "TRY",
		Timezone:      "Europe/Istanbul",
		MarketplaceId: "A33AVAJ2PDY3EV",
	},
	"amazon-br": {
		CurrencyCode:  "BRL",
		Timezone:      "America/Sao_Paulo",
		MarketplaceId: "A2Q3Y263D00KWC",
	},
	"amazon-ca": {
		CurrencyCode:  "CAD",
		Timezone:      "America/Toronto",
		MarketplaceId: "A2EUQ1WTGCTBG2",
	},
	"amazon-mx": {
		CurrencyCode:  "MXN",
		Timezone:      "America/Mexico_City",
		MarketplaceId: "A1AM78C64UM0Y8",
	},
	"amazon-us": {
		CurrencyCode:  "USD",
		Timezone:      "America/Los_Angeles",
		MarketplaceId: "ATVPDKIKX0DER",
	},
	"amazon-sg": {
		CurrencyCode:  "SGD",
		Timezone:      "Asia/Singapore",
		MarketplaceId: "A19VAU5U5O7RUS",
	},
	"amazon-au": {
		CurrencyCode:  "AUD",
		Timezone:      "Australia/Sydney",
		MarketplaceId: "A39IBJ37TRP1C6",
	},
	"amazon-jp": {
		CurrencyCode:  "JPY",
		Timezone:      "Asia/Tokyo",
		MarketplaceId: "A1VC38T7YXB528",
	},
}

func (o Channel) GetCurrencyCode() string {
	return ChannelMap[string(o)].CurrencyCode
}

func (o Channel) GetTimezone() string {
	return ChannelMap[string(o)].Timezone
}

func (o Channel) GetMarketplaceId() string {
	return ChannelMap[string(o)].MarketplaceId
}
