package model

// Fees between acquirer and issuer.
type DetailedAmount11 struct {

	// Type or class of amount.
	Type *TypeOfAmount7Code `xml:"Tp"`

	// Additional information to specify the type of amount.
	AdditionalType *Max35Text `xml:"AddtlTp,omitempty"`

	// Amount value.
	Amount *AmountAndDirection41 `xml:"Amt"`

	// Original value of the amount.
	OriginalAmount *AmountAndDirection41 `xml:"OrgnlAmt,omitempty"`
}

func (d *DetailedAmount11) SetType(value string) {
	d.Type = (*TypeOfAmount7Code)(&value)
}

func (d *DetailedAmount11) SetAdditionalType(value string) {
	d.AdditionalType = (*Max35Text)(&value)
}

func (d *DetailedAmount11) AddAmount() *AmountAndDirection41 {
	d.Amount = new(AmountAndDirection41)
	return d.Amount
}

func (d *DetailedAmount11) AddOriginalAmount() *AmountAndDirection41 {
	d.OriginalAmount = new(AmountAndDirection41)
	return d.OriginalAmount
}
