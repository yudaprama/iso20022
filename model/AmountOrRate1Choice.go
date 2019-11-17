package model

// Choice between an amount or a rate.
type AmountOrRate1Choice struct {

	// Amount expressed as an amount of money.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Amount expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (a *AmountOrRate1Choice) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountOrRate1Choice) SetRate(value string) {
	a.Rate = (*PercentageRate)(&value)
}
