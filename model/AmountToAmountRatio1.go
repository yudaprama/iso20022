package model

// Ratio expressed as a quotient of amounts.
type AmountToAmountRatio1 struct {

	// Numerator of the quotient of amounts.
	Amount1 *ActiveCurrencyAndAmount `xml:"Amt1"`

	// Denominator of the quotient of amounts
	Amount2 *ActiveCurrencyAndAmount `xml:"Amt2"`
}

func (a *AmountToAmountRatio1) SetAmount1(value, currency string) {
	a.Amount1 = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountToAmountRatio1) SetAmount2(value, currency string) {
	a.Amount2 = NewActiveCurrencyAndAmount(value, currency)
}
