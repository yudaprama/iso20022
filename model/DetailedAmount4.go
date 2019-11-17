package model

// Detailed amounts associated with the total amount of transaction.
type DetailedAmount4 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Short description of the amount to display or print.
	Label *Max140Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount4) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount4) SetLabel(value string) {
	d.Label = (*Max140Text)(&value)
}
