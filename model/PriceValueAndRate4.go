package model

// Price and rate.
type PriceValueAndRate4 struct {

	// Price expressed as a currency and amount.
	Value *PriceAndDirection1 `xml:"Val,omitempty"`

	// Price expressed as a percentage.
	Rate *PercentageRate `xml:"Rate,omitempty"`
}

func (p *PriceValueAndRate4) AddValue() *PriceAndDirection1 {
	p.Value = new(PriceAndDirection1)
	return p.Value
}

func (p *PriceValueAndRate4) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}
