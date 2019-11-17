package model

// Choice between an amount or rate.
type AmountOrRate3Choice struct {

	// Amount of money.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Rate used to calculate the amount
	Rate *PercentageRate `xml:"Rate"`
}

func (a *AmountOrRate3Choice) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (a *AmountOrRate3Choice) SetRate(value string) {
	a.Rate = (*PercentageRate)(&value)
}
