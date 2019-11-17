package model

// Choice between an amount or a rate.
type InterestRateUsedForPaymentFormat9Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for interest rate used for payment.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus32 `xml:"RateTpAndAmtAndRateSts"`
}

func (i *InterestRateUsedForPaymentFormat9Choice) SetRate(value string) {
	i.Rate = (*PercentageRate)(&value)
}

func (i *InterestRateUsedForPaymentFormat9Choice) SetAmount(value, currency string) {
	i.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (i *InterestRateUsedForPaymentFormat9Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus32 {
	i.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus32)
	return i.RateTypeAndAmountAndRateStatus
}
