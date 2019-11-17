package model

// Choice between an amount or a rate or an unspecified rate.
type TaxCreditRateFormat5Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the tax credit rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus5 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (t *TaxCreditRateFormat5Choice) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCreditRateFormat5Choice) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCreditRateFormat5Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus5 {
	t.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus5)
	return t.RateTypeAndAmountAndRateStatus
}

func (t *TaxCreditRateFormat5Choice) SetNotSpecifiedRate(value string) {
	t.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
