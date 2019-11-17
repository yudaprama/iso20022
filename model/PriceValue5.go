package model

// Value given to a price.
type PriceValue5 struct {

	// Price expressed as a currency and value.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (p *PriceValue5) SetAmount(value, currency string) {
	p.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}
