package model

// Choice between an amount or a rate or an unspecified rate.
type InterestRateUsedForPaymentFormat10Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for interest rate used for payment.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus32 `xml:"RateTpAndAmtAndRateSts"`

	// Value of the rate not specified.
	NotSpecifiedRate *RateType13Code `xml:"NotSpcfdRate"`
}

func (i *InterestRateUsedForPaymentFormat10Choice) SetRate(value string) {
	i.Rate = (*PercentageRate)(&value)
}

func (i *InterestRateUsedForPaymentFormat10Choice) SetAmount(value, currency string) {
	i.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (i *InterestRateUsedForPaymentFormat10Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus32 {
	i.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus32)
	return i.RateTypeAndAmountAndRateStatus
}

func (i *InterestRateUsedForPaymentFormat10Choice) SetNotSpecifiedRate(value string) {
	i.NotSpecifiedRate = (*RateType13Code)(&value)
}
