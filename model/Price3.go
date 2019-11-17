package model

// Type and information about a price.
type Price3 struct {

	// Specification of the price type.
	Type *YieldedOrValueType1Choice `xml:"Tp"`

	// Value of the price, eg, as a currency and value.
	Value *PriceRateOrAmount1Choice `xml:"Val"`
}

func (p *Price3) AddType() *YieldedOrValueType1Choice {
	p.Type = new(YieldedOrValueType1Choice)
	return p.Type
}

func (p *Price3) AddValue() *PriceRateOrAmount1Choice {
	p.Value = new(PriceRateOrAmount1Choice)
	return p.Value
}
