package model

// Factors used in the calculation of the pay in schedule.
type PayInFactors1 struct {

	// Maximum allowed sum of short positions in all currencies, converted to base currency, during settlement.
	AggregateShortPositionLimit *ActiveCurrencyAndAmount `xml:"AggtShrtPosLmt"`

	// Currency specific pay-in factors.
	CurrencyFactors []*CurrencyFactors1 `xml:"CcyFctrs"`
}

func (p *PayInFactors1) SetAggregateShortPositionLimit(value, currency string) {
	p.AggregateShortPositionLimit = NewActiveCurrencyAndAmount(value, currency)
}

func (p *PayInFactors1) AddCurrencyFactors() *CurrencyFactors1 {
	newValue := new(CurrencyFactors1)
	p.CurrencyFactors = append(p.CurrencyFactors, newValue)
	return newValue
}
