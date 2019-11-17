package model

// Withdrawal fees, accepted by the customer.
type DetailedAmount13 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Short description of the amount to display or print.
	Label *Max70Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount13) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount13) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount13) SetLabel(value string) {
	d.Label = (*Max70Text)(&value)
}
