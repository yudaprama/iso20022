package model

// Choice of formats for the price.
type PriceRateOrAmountOrUnknownChoice struct {

	// Price expressed as a rate, ie, percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Price expressed as a currency and value.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Indicates whether the market price is unknown.
	UnknownIndicator *YesNoIndicator `xml:"UknwnInd"`
}

func (p *PriceRateOrAmountOrUnknownChoice) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}

func (p *PriceRateOrAmountOrUnknownChoice) SetAmount(value, currency string) {
	p.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceRateOrAmountOrUnknownChoice) SetUnknownIndicator(value string) {
	p.UnknownIndicator = (*YesNoIndicator)(&value)
}
