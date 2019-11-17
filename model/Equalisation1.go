package model

// Part of an investor's subscription amount that is held by the fund in order to pay incentive / performance fees at the end of the fiscal year.
type Equalisation1 struct {

	// Amount of money resulting from the calculation of the equalisation.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Rate used for calculation of the equalisation.
	Rate *PercentageRate `xml:"Rate,omitempty"`
}

func (e *Equalisation1) SetAmount(value, currency string) {
	e.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (e *Equalisation1) SetRate(value string) {
	e.Rate = (*PercentageRate)(&value)
}
