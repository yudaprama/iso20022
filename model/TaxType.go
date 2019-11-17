package model

// Information on the type of tax.
type TaxType struct {

	// Description of the tax that is being paid, including specific representation required by taxing authority.
	CategoryDescription *Max35Text `xml:"CtgyDesc,omitempty"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate,omitempty"`

	// Amount of money on which the tax is based.
	TaxableBaseAmount *CurrencyAndAmount `xml:"TaxblBaseAmt,omitempty"`

	// Amount of money resulting from the calculation of the tax.
	Amount *CurrencyAndAmount `xml:"Amt,omitempty"`
}

func (t *TaxType) SetCategoryDescription(value string) {
	t.CategoryDescription = (*Max35Text)(&value)
}

func (t *TaxType) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}

func (t *TaxType) SetTaxableBaseAmount(value, currency string) {
	t.TaxableBaseAmount = NewCurrencyAndAmount(value, currency)
}

func (t *TaxType) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}
