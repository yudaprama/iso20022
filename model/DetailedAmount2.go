package model

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount2 struct {

	// Type of amount.
	Type *TypeOfAmount2Code `xml:"Tp"`

	// Amount value.
	Value *ImpliedCurrencyAndAmount `xml:"Val"`

	// Short description of the amount.
	Label *Max35Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount2) SetType(value string) {
	d.Type = (*TypeOfAmount2Code)(&value)
}

func (d *DetailedAmount2) SetValue(value, currency string) {
	d.Value = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount2) SetLabel(value string) {
	d.Label = (*Max35Text)(&value)
}
