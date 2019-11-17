package model

// Specifies a price expressed as a rate.
type PriceRate1 struct {

	// Type of rate, eg, yield.
	RateType *PriceRateType3FormatChoice `xml:"RateTp"`

	// Price expressed as a rate, ie, percentage.
	Rate *PercentageRate `xml:"Rate"`
}

func (p *PriceRate1) AddRateType() *PriceRateType3FormatChoice {
	p.RateType = new(PriceRateType3FormatChoice)
	return p.RateType
}

func (p *PriceRate1) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}
