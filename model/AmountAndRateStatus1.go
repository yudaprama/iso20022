package model

// Specifies an amount and a rate status.
type AmountAndRateStatus1 struct {

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Value expressed as a rate status.
	RateStatus *RateStatus1Code `xml:"RateSts"`
}

func (a *AmountAndRateStatus1) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountAndRateStatus1) SetRateStatus(value string) {
	a.RateStatus = (*RateStatus1Code)(&value)
}
