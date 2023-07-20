package amazon

import (
	"github.com/amazinsellers/ecommerce-models/amazon"
	"testing"
)

func TestChannelGetCurrencyCodeReturnsCorrectCurrencyCode(t *testing.T) {
	currency := (amazon.Channel("amazon-us")).GetCurrencyCode()

	if currency != "USD" {
		t.Error("currency for amazon-us is not returning USD, rather as " + currency)
	}
}

func TestChannelGetTimezoneReturnsCorrectTimezone(t *testing.T) {
	tz := (amazon.Channel("amazon-in")).GetTimezone()

	if tz != "Asia/Kolkata" {
		t.Error("timezone for amazon-in is not returning \"Asia/Kolkata\", rather as " + tz)
	}
}
