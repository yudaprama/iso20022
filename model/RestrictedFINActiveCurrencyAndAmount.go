package model

// A number of monetary units specified in an active currency where the unit of currency is explicit and compliant with ISO 4217.
type RestrictedFINActiveCurrencyAndAmount struct {
	Value    string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewRestrictedFINActiveCurrencyAndAmount(value, currency string) *RestrictedFINActiveCurrencyAndAmount {
	return &RestrictedFINActiveCurrencyAndAmount{Value: value, Currency: currency}
}
