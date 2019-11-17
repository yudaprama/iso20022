package model

// Choice of formats for the price.
type PriceRateOrAmountOrUnknown1Choice struct {

	// Price expressed as a rate, ie, percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Price expressed as a currency and value.
	Amount *RestrictedFINActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Indicates whether the market price is unknown.
	UnknownIndicator *YesNoIndicator `xml:"UknwnInd"`
}

func (p *PriceRateOrAmountOrUnknown1Choice) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}

func (p *PriceRateOrAmountOrUnknown1Choice) SetAmount(value, currency string) {
	p.Amount = NewRestrictedFINActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceRateOrAmountOrUnknown1Choice) SetUnknownIndicator(value string) {
	p.UnknownIndicator = (*YesNoIndicator)(&value)
}
