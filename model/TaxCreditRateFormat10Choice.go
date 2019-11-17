package model

// Choice between an amount or a rate or an unspecified rate.
type TaxCreditRateFormat10Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the tax credit rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus34 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateValueType7Code `xml:"NotSpcfdRate"`
}

func (t *TaxCreditRateFormat10Choice) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCreditRateFormat10Choice) SetAmount(value, currency string) {
	t.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCreditRateFormat10Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus34 {
	t.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus34)
	return t.RateTypeAndAmountAndRateStatus
}

func (t *TaxCreditRateFormat10Choice) SetNotSpecifiedRate(value string) {
	t.NotSpecifiedRate = (*RateValueType7Code)(&value)
}
