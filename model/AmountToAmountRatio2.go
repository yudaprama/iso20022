package model

// Ratio expressed as a quotient of amounts.
type AmountToAmountRatio2 struct {

	// Numerator of the quotient of amounts.
	Amount1 *ActiveCurrencyAnd13DecimalAmount `xml:"Amt1"`

	// Denominator of the quotient of amounts
	Amount2 *ActiveCurrencyAnd13DecimalAmount `xml:"Amt2"`
}

func (a *AmountToAmountRatio2) SetAmount1(value, currency string) {
	a.Amount1 = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountToAmountRatio2) SetAmount2(value, currency string) {
	a.Amount2 = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}
