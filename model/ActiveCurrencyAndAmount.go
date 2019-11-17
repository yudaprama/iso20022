package model

// A number of monetary units specified in an active currency where the unit of currency is explicit and compliant with ISO 4217.
type ActiveCurrencyAndAmount struct {
	Value    string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewActiveCurrencyAndAmount(value, currency string) *ActiveCurrencyAndAmount {
	return &ActiveCurrencyAndAmount{Value: value, Currency: currency}
}
