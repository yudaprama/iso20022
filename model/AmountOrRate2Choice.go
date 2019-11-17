package model

// Choice between an amount or a rate.
type AmountOrRate2Choice struct {

	// Amount of money in the cash entry.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Amount expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (a *AmountOrRate2Choice) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountOrRate2Choice) SetRate(value string) {
	a.Rate = (*PercentageRate)(&value)
}
