package model

// Choice between an amount or a rate or an unspecified rate.
type TaxCreditRateFormat1Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the tax credit rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus5 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType6Code `xml:"NotSpcfdRate"`
}

func (t *TaxCreditRateFormat1Choice) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCreditRateFormat1Choice) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCreditRateFormat1Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus5 {
	t.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus5)
	return t.RateTypeAndAmountAndRateStatus
}

func (t *TaxCreditRateFormat1Choice) SetNotSpecifiedRate(value string) {
	t.NotSpecifiedRate = (*RateValueType6Code)(&value)
}
