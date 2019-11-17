package model

// Number of monetary units specified in a currency, where the unit of currency is explicit and compliant with ISO 4217.  The decimal separator is a dot.
// Note: A zero amount is considered a positive amount.
type CurrencyAndAmount struct {
	Value    string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewCurrencyAndAmount(value, currency string) *CurrencyAndAmount {
	return &CurrencyAndAmount{Value: value, Currency: currency}
}
