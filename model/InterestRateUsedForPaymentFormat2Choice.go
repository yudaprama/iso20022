package model

// Choice between an amount or a rate.
type InterestRateUsedForPaymentFormat2Choice struct {

	// Value is expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Number of monetary units specified in a currency.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for interest rate used for payment.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus3 `xml:"RateTpAndAmtAndRateSts"`
}

func (i *InterestRateUsedForPaymentFormat2Choice) SetRate(value string) {
	i.Rate = (*PercentageRate)(&value)
}

func (i *InterestRateUsedForPaymentFormat2Choice) SetAmount(value, currency string) {
	i.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (i *InterestRateUsedForPaymentFormat2Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus3 {
	i.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus3)
	return i.RateTypeAndAmountAndRateStatus
}
