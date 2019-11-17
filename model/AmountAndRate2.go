package model

// Change amount and rate.
type AmountAndRate2 struct {

	// Amount expressed as an amount of money.
	Amount *AmountAndDirection30 `xml:"Amt,omitempty"`

	// Amount expressed as a rate.
	Rate *PercentageRate `xml:"Rate,omitempty"`
}

func (a *AmountAndRate2) AddAmount() *AmountAndDirection30 {
	a.Amount = new(AmountAndDirection30)
	return a.Amount
}

func (a *AmountAndRate2) SetRate(value string) {
	a.Rate = (*PercentageRate)(&value)
}
