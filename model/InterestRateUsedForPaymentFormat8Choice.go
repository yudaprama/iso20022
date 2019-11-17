package model

// Choice between an amount or a rate or an unspecified rate.
type InterestRateUsedForPaymentFormat8Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for interest rate used for payment.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus24 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType13Code `xml:"NotSpcfdRate"`
}

func (i *InterestRateUsedForPaymentFormat8Choice) SetRate(value string) {
	i.Rate = (*PercentageRate)(&value)
}

func (i *InterestRateUsedForPaymentFormat8Choice) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (i *InterestRateUsedForPaymentFormat8Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus24 {
	i.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus24)
	return i.RateTypeAndAmountAndRateStatus
}

func (i *InterestRateUsedForPaymentFormat8Choice) SetNotSpecifiedRate(value string) {
	i.NotSpecifiedRate = (*RateType13Code)(&value)
}
