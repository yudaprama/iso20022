package model

// Limit of amounts for the customer.
type ATMTransactionAmounts6 struct {

	// Default currency of the limits.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Maximum amount allowed for a transaction in the service.
	MaximumPossibleAmount *ImpliedCurrencyAndAmount `xml:"MaxPssblAmt,omitempty"`

	// Minimum amount allowed for a transaction in the service.
	MinimumPossibleAmount *ImpliedCurrencyAndAmount `xml:"MinPssblAmt,omitempty"`

	// Additional amount that may be displayed to the customer, for instance the daily limit or the daily balance for the service.
	AdditionalAmount []*ATMTransactionAmounts7 `xml:"AddtlAmt,omitempty"`
}

func (a *ATMTransactionAmounts6) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts6) SetMaximumPossibleAmount(value, currency string) {
	a.MaximumPossibleAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts6) SetMinimumPossibleAmount(value, currency string) {
	a.MinimumPossibleAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts6) AddAdditionalAmount() *ATMTransactionAmounts7 {
	newValue := new(ATMTransactionAmounts7)
	a.AdditionalAmount = append(a.AdditionalAmount, newValue)
	return newValue
}
