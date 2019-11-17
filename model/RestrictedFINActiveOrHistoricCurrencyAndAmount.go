package model

// Number of monetary units specified in an active or a historic currency where the unit of currency is explicit and compliant with ISO 4217.
type RestrictedFINActiveOrHistoricCurrencyAndAmount struct {
	Value    string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewRestrictedFINActiveOrHistoricCurrencyAndAmount(value, currency string) *RestrictedFINActiveOrHistoricCurrencyAndAmount {
	return &RestrictedFINActiveOrHistoricCurrencyAndAmount{Value: value, Currency: currency}
}
