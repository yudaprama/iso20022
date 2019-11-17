package model

// Ratio expressed as a quotient of amounts.
type AmountToAmountRatio3 struct {

	// Numerator of the quotient of amounts.
	Amount1 *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt1"`

	// Denominator of the quotient of amounts
	Amount2 *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt2"`
}

func (a *AmountToAmountRatio3) SetAmount1(value, currency string) {
	a.Amount1 = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountToAmountRatio3) SetAmount2(value, currency string) {
	a.Amount2 = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}
