package model

// Choice between amount expressed as an absolute value or as a percentage rate related to another reference amount.
type FinancingRateOrAmountChoice struct {

	// Amount expressed as an absolute value.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`

	// Amount expressed as a percentage rate.
	Rate *PercentageRate `xml:"Rate"`
}

func (f *FinancingRateOrAmountChoice) SetAmount(value, currency string) {
	f.Amount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FinancingRateOrAmountChoice) SetRate(value string) {
	f.Rate = (*PercentageRate)(&value)
}
