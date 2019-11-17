package model

// Defined variation amount and balance.
type UndertakingAmount4 struct {

	// Variation amount and currency.
	VariationAmount *ActiveCurrencyAndAmount `xml:"VartnAmt"`

	// Calculated undertaking available balance amount resulting from the application of the variation amount.
	BalanceAmount *ActiveCurrencyAndAmount `xml:"BalAmt,omitempty"`
}

func (u *UndertakingAmount4) SetVariationAmount(value, currency string) {
	u.VariationAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (u *UndertakingAmount4) SetBalanceAmount(value, currency string) {
	u.BalanceAmount = NewActiveCurrencyAndAmount(value, currency)
}
