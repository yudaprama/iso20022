package model

// Signed amount.
type AmountAndDirection43 struct {

	// Amount value.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Indicates that the amount value is positive or negative.
	Sign *PlusOrMinusIndicator `xml:"Sgn,omitempty"`

	// Date of the amount.
	Date *ISODate `xml:"Dt,omitempty"`
}

func (a *AmountAndDirection43) SetAmount(value, currency string) {
	a.Amount = NewCurrencyAndAmount(value, currency)
}

func (a *AmountAndDirection43) SetSign(value string) {
	a.Sign = (*PlusOrMinusIndicator)(&value)
}

func (a *AmountAndDirection43) SetDate(value string) {
	a.Date = (*ISODate)(&value)
}
