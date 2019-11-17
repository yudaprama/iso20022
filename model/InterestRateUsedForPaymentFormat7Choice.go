package model

// Choice between an amount or a rate.
type InterestRateUsedForPaymentFormat7Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for interest rate used for payment.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus24 `xml:"RateTpAndAmtAndRateSts"`
}

func (i *InterestRateUsedForPaymentFormat7Choice) SetRate(value string) {
	i.Rate = (*PercentageRate)(&value)
}

func (i *InterestRateUsedForPaymentFormat7Choice) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (i *InterestRateUsedForPaymentFormat7Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus24 {
	i.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus24)
	return i.RateTypeAndAmountAndRateStatus
}
