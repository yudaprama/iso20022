package model

// Choice between an amount or an unspecified rate.
type TaxCreditRateFormat9Choice struct {

	// Value expressed as a rate.
	Rate *PercentageRate `xml:"Rate"`

	// Value expressed as an amount.
	Amount *RestrictedFINActiveCurrencyAnd13DecimalAmount `xml:"Amt"`

	// Specifies different formats for the tax credit rate.
	RateTypeAndAmountAndRateStatus *RateTypeAndAmountAndStatus34 `xml:"RateTpAndAmtAndRateSts"`
}

func (t *TaxCreditRateFormat9Choice) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxCreditRateFormat9Choice) SetAmount(value, currency string) {
	t.Amount = NewRestrictedFINActiveCurrencyAnd13DecimalAmount(value, currency)
}

func (t *TaxCreditRateFormat9Choice) AddRateTypeAndAmountAndRateStatus() *RateTypeAndAmountAndStatus34 {
	t.RateTypeAndAmountAndRateStatus = new(RateTypeAndAmountAndStatus34)
	return t.RateTypeAndAmountAndRateStatus
}
