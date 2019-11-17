package model

// Withdrawal fees, accepted by the customer.
type DetailedAmount18 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// True if amount charged to the source account.
	ChargeAccountTo *TrueFalseIndicator `xml:"ChrgAcctTo,omitempty"`

	// Short description of the amount to display or print.
	Label *Max70Text `xml:"Labl,omitempty"`
}

func (d *DetailedAmount18) SetAmount(value, currency string) {
	d.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (d *DetailedAmount18) SetCurrency(value string) {
	d.Currency = (*ActiveCurrencyCode)(&value)
}

func (d *DetailedAmount18) SetChargeAccountTo(value string) {
	d.ChargeAccountTo = (*TrueFalseIndicator)(&value)
}

func (d *DetailedAmount18) SetLabel(value string) {
	d.Label = (*Max70Text)(&value)
}
