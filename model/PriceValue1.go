package model

// Value given to a price.
type PriceValue1 struct {

	// Price expressed as a currency and value.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (p *PriceValue1) SetAmount(value, currency string) {
	p.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
