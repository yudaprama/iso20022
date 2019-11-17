package model

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount9 struct {

	// Type or class of amount.
	Type *TypeOfAmount5Code `xml:"Tp"`

	// Additional information to specify the type of amount.
	AdditionalType *Max35Text `xml:"AddtlTp,omitempty"`

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Short description of the amount to provide to the cardholder.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount9) SetType(value string) {
	d.Type = (*TypeOfAmount5Code)(&value)
}

func (d *DetailedAmount9) SetAdditionalType(value string) {
	d.AdditionalType = (*Max35Text)(&value)
}

func (d *DetailedAmount9) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount9) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}
