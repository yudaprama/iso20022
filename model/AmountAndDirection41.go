package model

// Signed amount.
type AmountAndDirection41 struct {

	// Amount value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`
}

func (a *AmountAndDirection41) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection41) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}
