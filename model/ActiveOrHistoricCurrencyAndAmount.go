package model

// A number of monetary units specified in an active or a historic currency where the unit of currency is explicit and compliant with ISO 4217.
type ActiveOrHistoricCurrencyAndAmount struct {
	Value    string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewActiveOrHistoricCurrencyAndAmount(value, currency string) *ActiveOrHistoricCurrencyAndAmount {
	return &ActiveOrHistoricCurrencyAndAmount{Value: value, Currency: currency}
}
