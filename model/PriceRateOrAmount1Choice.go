package model

// Choice of formats for the price.
type PriceRateOrAmount1Choice struct {

	// Price expressed as a rate, ie, percentage.
	Rate *PercentageRate `xml:"Rate"`

	// Price expressed as a currency and value.
	Amount *RestrictedFINActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt"`
}

func (p *PriceRateOrAmount1Choice) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}

func (p *PriceRateOrAmount1Choice) SetAmount(value, currency string) {
	p.Amount = NewRestrictedFINActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}
