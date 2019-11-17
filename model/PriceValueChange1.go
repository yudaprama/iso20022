package model

// Value given to a positive or negative price change.
type PriceValueChange1 struct {

	// Amount by which the price has changed.
	Amount *ActiveOrHistoricCurrencyAnd13DecimalAmount `xml:"Amt,omitempty"`

	// Indicates a positive or negative amount change.
	AmountSign *PlusOrMinusIndicator `xml:"AmtSgn,omitempty"`

	// Rate by which the price has changed.
	Rate *PercentageRate `xml:"Rate,omitempty"`
}

func (p *PriceValueChange1) SetAmount(value, currency string) {
	p.Amount = NewActiveOrHistoricCurrencyAnd13DecimalAmount(value, currency)
}

func (p *PriceValueChange1) SetAmountSign(value string) {
	p.AmountSign = (*PlusOrMinusIndicator)(&value)
}

func (p *PriceValueChange1) SetRate(value string) {
	p.Rate = (*PercentageRate)(&value)
}
