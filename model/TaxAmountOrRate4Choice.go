package model

// Choice between formats for the specification of a tax.
type TaxAmountOrRate4Choice struct {

	// Amount of money resulting from the calculation of the tax.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate"`
}

func (t *TaxAmountOrRate4Choice) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (t *TaxAmountOrRate4Choice) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}
