package model

// Amount to be authorised by the issuer.
type AmountAndCurrency1 struct {

	// Amount value.
	Amount *ImpliedCurrencyAndAmount `xml:"Amt"`

	// Currency of the amount.
	Currency *ActiveCurrencyCode `xml:"Ccy,omitempty"`
}

func (a *AmountAndCurrency1) SetAmount(value, currency string) {
	a.Amount = NewImpliedCurrencyAndAmount(value, currency)
}

func (a *AmountAndCurrency1) SetCurrency(value string) {
	a.Currency = (*ActiveCurrencyCode)(&value)
}
