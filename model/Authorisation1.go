package model

// Autorisation of the mandate holder.
type Authorisation1 struct {

	// Minimum amount per transaction allowed by the mandate.
	MinimumAmountPerTransaction *ActiveCurrencyAndAmount `xml:"MinAmtPerTx"`

	// Maximum amount per transaction allowed by the mandate.
	MaximumAmountPerTransaction *ActiveCurrencyAndAmount `xml:"MaxAmtPerTx"`

	// Maximum amount allowed over a specific period of time.
	MaximumAmountByPeriod []*MaximumAmountByPeriod1 `xml:"MaxAmtByPrd,omitempty"`
}

func (a *Authorisation1) SetMinimumAmountPerTransaction(value, currency string) {
	a.MinimumAmountPerTransaction = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Authorisation1) SetMaximumAmountPerTransaction(value, currency string) {
	a.MaximumAmountPerTransaction = NewActiveCurrencyAndAmount(value, currency)
}

func (a *Authorisation1) AddMaximumAmountByPeriod() *MaximumAmountByPeriod1 {
	newValue := new(MaximumAmountByPeriod1)
	a.MaximumAmountByPeriod = append(a.MaximumAmountByPeriod, newValue)
	return newValue
}
