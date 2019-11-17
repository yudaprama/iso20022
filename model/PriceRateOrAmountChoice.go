package model

// Choice of formats for the price.
type PriceRateOrAmountChoice struct {

	// Price expressed as a rate, that is percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Price expressed as a currency and value.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (p *PriceRateOrAmountChoice) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}

func (p *PriceRateOrAmountChoice) SetAmount(value, currency string) {
	p.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}
