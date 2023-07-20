package amazon

type Channel string

type ChannelProp struct {
	CurrencyCode string
	Timezone     string
}

var ChannelMap = map[string]ChannelProp{
	"amazon-ae": {
		CurrencyCode: "AED",
		Timezone:     "Asia/Dubai",
	},
	"amazon-de": {
		CurrencyCode: "EUR",
		Timezone:     "Europe/Berlin",
	},
	"amazon-eg": {
		CurrencyCode: "EGP",
		Timezone:     "Africa/Cairo",
	},
	"amazon-es": {
		CurrencyCode: "EUR",
		Timezone:     "Europe/Madrid",
	},
	"amazon-fr": {
		CurrencyCode: "EUR",
		Timezone:     "Europe/Paris",
	},
	"amazon-gb": {
		CurrencyCode: "GBP",
		Timezone:     "Europe/London",
	},
	"amazon-in": {
		CurrencyCode: "INR",
		Timezone:     "Asia/Kolkata",
	},
	"amazon-nl": {
		CurrencyCode: "EUR",
		Timezone:     "Europe/Amsterdam",
	},
	"amazon-pl": {
		CurrencyCode: "PLN",
		Timezone:     "Europe/Warsaw",
	},
	"amazon-sa": {
		CurrencyCode: "EUR",
		Timezone:     "Asia/Riyadh",
	},
	"amazon-se": {
		CurrencyCode: "EUR",
		Timezone:     "Europe/Stockholm",
	},
	"amazon-tr": {
		CurrencyCode: "TRY",
		Timezone:     "Europe/Istanbul",
	},
	"amazon-br": {
		CurrencyCode: "BRL",
		Timezone:     "America/Sao_Paulo",
	},
	"amazon-ca": {
		CurrencyCode: "CAD",
		Timezone:     "America/Toronto",
	},
	"amazon-mx": {
		CurrencyCode: "MXN",
		Timezone:     "America/Mexico_City",
	},
	"amazon-us": {
		CurrencyCode: "USD",
		Timezone:     "America/Los_Angeles",
	},
	"amazon-sg": {
		CurrencyCode: "SGD",
		Timezone:     "Asia/Singapore",
	},
	"amazon-au": {
		CurrencyCode: "AUD",
		Timezone:     "Australia/Sydney",
	},
	"amazon-jp": {
		CurrencyCode: "JPY",
		Timezone:     "Asia/Tokyo",
	},
}

func (o Channel) GetCurrencyCode() string {
	return ChannelMap[string(o)].CurrencyCode
}

func (o Channel) GetTimezone() string {
	return ChannelMap[string(o)].Timezone
}
