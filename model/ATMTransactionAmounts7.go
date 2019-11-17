package model

// Additional amount that may be displayed to the customer, for instance the daily limit or the daily balance for the service.
type ATMTransactionAmounts7 struct {

	// Type of amount.
	Type *Max35Text `xml:"Tp"`

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Description of the amount that may be provided to the customer.
	Label *Max70Text `xml:"Labl,omitempty"`
}

func (a *ATMTransactionAmounts7) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *ATMTransactionAmounts7) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts7) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts7) SetLabel(value string) {
	a.Label = (*Max70Text)(&value)
}
