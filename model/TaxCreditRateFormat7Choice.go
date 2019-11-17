package model

// Choice between an amount or an unspecified rate.
type TaxCreditRateFormat7Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value expressed as an amount.
	Amount *ActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the tax credit rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus27 `xml:"RateTpAndAmtAndRateSts"`
}

func (t *TaxCreditRateFormat7Choice) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCreditRateFormat7Choice) SetAmount(value, currency string) {
	t.Amount = NewActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCreditRateFormat7Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus27 {
	t.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus27)
	return t.RateTypeAndAmountAndRateStatus
}
