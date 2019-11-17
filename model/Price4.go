package model

// Provides the value and optionaly the type of price.
type Price4 struct {

	// Value of the price.
	Value *PriceRateOrAmountChoice `xml:"Val"`

	// Specification of the price type.
	Type *PriceValueType7Code `xml:"Tp,omitempty"`
}

func (p *Price4) AddValue() *PriceRateOrAmountChoice {
	p.Value = new(PriceRateOrAmountChoice)
	return p.Value
}

func (p *Price4) SetType(value string) {
	p.Type = (*PriceValueType7Code)(&value)
}
