package model

// Choice between an amount or an unspecified rate.
type TaxCreditRateFormat2Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the tax credit rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus5 `xml:"RateTpAndAmtAndRateSts"`
}

func (t *TaxCreditRateFormat2Choice) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCreditRateFormat2Choice) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCreditRateFormat2Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus5 {
	t.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus5)
	return t.RateTypeAndAmountAndRateStatus
}
