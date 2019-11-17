package model

// A number of monetary units specified in an active or a historic currency where the unit of currency is explicit and compliant with ISO 4217. The number of fractional digits (or minor unit of currency) is not checked as per ISO 4217: It must be lesser than or equal to 13.
// Note: The decimal separator is a dot.
type ActiveOrHistoricCurrencyAnd13DecimalAmount struct {
	Value    string `xml:",chardata"`
	Currency string `xml:"Ccy,attr"`
}

func NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency string) *ActiveOrHistoricCurrencyAnd13DecimalAmount {
	return &ActiveOrHistoricCurrencyAnd13DecimalAmount{Value: value, Currency: currency}
}
