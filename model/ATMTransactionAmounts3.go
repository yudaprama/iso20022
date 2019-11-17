package model

// Limit of amounts for the customer.
type ATMTransactionAmounts3 struct {

	// Type of limit.
	Type *Max35Text `xml:"Tp"`

	// Label of the limit to display or print.
	Label *Max35Text `xml:"Labl,omitempty"`

	// Currency of the limit amount.
	Currency *ActiveCurrencyCode `xml:"Ccy"`

	// Minimum amount value in the currency of the limit.
	MinimumAmount *ImpliedCurrencyAndAmount `xml:"MinAmt,omitempty"`

	// Maximum amount value in the currency of the limit.
	MaximumAmount *ImpliedCurrencyAndAmount `xml:"MaxAmt,omitempty"`
}

func (a *ATMTransactionAmounts3) SetType(value string) {
	a.Type = (*Max35Text)(&value)
}

func (a *ATMTransactionAmounts3) SetLabel(value string) {
	a.Label = (*Max35Text)(&value)
}

func (a *ATMTransactionAmounts3) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts3) SetMinimumAmount(value, currency string) {
	a.MinimumAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts3) SetMaximumAmount(value, currency string) {
	a.MaximumAmount = NewImpliedCurrencyAndAmount(value, currency)
}
