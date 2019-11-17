package model

// Choice between an amount or a rate or an unspecified rate.
type TaxCreditRateFormat8Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the tax credit rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus27 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (t *TaxCreditRateFormat8Choice) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCreditRateFormat8Choice) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCreditRateFormat8Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus27 {
	t.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus27)
	return t.RateTypeAndAmountAndRateStatus
}

func (t *TaxCreditRateFormat8Choice) SetNotSpecifiedRate(value string) {
	t.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
