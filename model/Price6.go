package model

// Provides the value, type and source of price.
type Price6 struct {

	// Value of the price expressed as a currency and value or as a rate.
	RateOrAmount *PriceRateOrAmountChoice `xml:"RateOrAmt"`

	// Specification of the price type.
	Type *TypeOfPrice13Code `xml:"Tp"`

	// Source for the price valuation.
	Source *PriceSource2Code `xml:"Src"`
}

func (p *Price6) AddRateOrAmount() *PriceRateOrAmountChoice {
	p.RateOrAmount = new(PriceRateOrAmountChoice)
	return p.RateOrAmount
}

func (p *Price6) SetType(value string) {
	p.Type = (*TypeOfPrice13Code)(&value)
}

func (p *Price6) SetSource(value string) {
	p.Source = (*PriceSource2Code)(&value)
}
