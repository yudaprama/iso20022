package model

// Limit of amounts for the customer.
type ATMTransactionAmounts8 struct {

	// Default currency of the limits.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`

	// Maximum amount allowed for a transaction in the service.
	MaximumPossibleAmount *ImpliedCurrencyAndAmount `xml:"MaxPssblAmt,omitempty"`

	// Minimum amount allowed for a transaction in the service.
	MinimumPossibleAmount *ImpliedCurrencyAndAmount `xml:"MinPssblAmt,omitempty"`

	// Additional amount that may be displayed to the customer, for instance the daily limit or the daily balance for the service.
	AdditionalAmount []*ATMTransactionAmounts7 `xml:"AddtlAmt,omitempty"`

	// Limit of deposited media for the customer.
	DepositLimits []*ATMTransactionAmounts9 `xml:"DpstLmts,omitempty"`
}

func (a *ATMTransactionAmounts8) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}

func (a *ATMTransactionAmounts8) SetMaximumPossibleAmount(value, currency string) {
	a.MaximumPossibleAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts8) SetMinimumPossibleAmount(value, currency string) {
	a.MinimumPossibleAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *ATMTransactionAmounts8) AddAdditionalAmount() *ATMTransactionAmounts7 {
	newValue := new(ATMTransactionAmounts7)
	a.AdditionalAmount = append(a.AdditionalAmount, newValue)
	return newValue
}

func (a *ATMTransactionAmounts8) AddDepositLimits() *ATMTransactionAmounts9 {
	newValue := new(ATMTransactionAmounts9)
	a.DepositLimits = append(a.DepositLimits, newValue)
	return newValue
}
