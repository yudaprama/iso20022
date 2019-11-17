package model

// Specifies the payment terms of the underlying transaction.
type EarlyPaymentsVAT1 struct {

	// Tax rate to be applied for early payment.
	TaxRate *PercentageRate `xml:"TaxRate"`

	// Type of tax applied.
	DiscountTaxType *Max4Text `xml:"DscntTaxTp"`

	// Early payment discount tax amount calculated using defined tax rate.
	DiscountTaxAmount *CurrencyAndAmount `xml:"DscntTaxAmt"`
}

func (e *EarlyPaymentsVAT1) SetTaxRate(value string) {
	e.TaxRate = (*PercentageRate)(&value)
}

func (e *EarlyPaymentsVAT1) SetDiscountTaxType(value string) {
	e.DiscountTaxType = (*Max4Text)(&value)
}

func (e *EarlyPaymentsVAT1) SetDiscountTaxAmount(value, currency string) {
	e.DiscountTaxAmount = NewCurrencyAndAmount(value, currency)
}
