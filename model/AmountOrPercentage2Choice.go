package model

// Choice between a set amount or a percentage.
type AmountOrPercentage2Choice struct {

	// Specifies that the payment is for a fixed amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Specifies that the payment conditions apply to a percentage of the amount due.
	Percentage *PercentageRate `xml:"Pctg"`
}

func (a *AmountOrPercentage2Choice) SetAmount(value, currency string) {
	a.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (a *AmountOrPercentage2Choice) SetPercentage(value string) {
	a.Percentage = (*PercentageRate)(&value)
}
