package model

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax13 struct {

	// Type of tax applied.
	Type *TaxType9Code `xml:"Tp"`

	// Specifies types of tax not present in a code list.
	OtherTaxType *Max35Text `xml:"OthrTaxTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *CurrencyAndAmount `xml:"Amt"`

	// Rate used to calculate the tax.
	Rate *PercentageRate `xml:"Rate"`
}

func (t *Tax13) SetType(value string) {
	t.Type = (*TaxType9Code)(&value)
}

func (t *Tax13) SetOtherTaxType(value string) {
	t.OtherTaxType = (*Max35Text)(&value)
}

func (t *Tax13) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}

func (t *Tax13) SetRate(value string) {
	t.Rate = (*PercentageRate)(&value)
}
