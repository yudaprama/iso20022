package model

// Price and an indication of whether it is a increase or a decrease.
type PriceAndDirection1 struct {

	// Currency and value.
	Value *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Val"`

	// Indicates that the value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`
}

func (p *PriceAndDirection1) SetValue(value, currency string) {
	p.Value = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceAndDirection1) SetSign(value string) {
	p.Sign = (*PlusOrMinusIndicator)(&value)
}
