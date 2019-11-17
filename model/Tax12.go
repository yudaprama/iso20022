package model

// Amount of money due to the government or tax authority, according to various pre-defined parameters such as thresholds or income.
type Tax12 struct {

	// Type of tax applied.
	Type *TaxType9Code `xml:"Tp"`

	// Specifies types of tax not present in a code list.
	OtherTaxType *Max35Text `xml:"OthrTaxTp"`

	// Amount of money resulting from the calculation of the tax.
	Amount *CurrencyAndAmount `xml:"Amt"`
}

func (t *Tax12) SetType(value string) {
	t.Type = (*TaxType9Code)(&value)
}

func (t *Tax12) SetOtherTaxType(value string) {
	t.OtherTaxType = (*Max35Text)(&value)
}

func (t *Tax12) SetAmount(value, currency string) {
	t.Amount = NewCurrencyAndAmount(value, currency)
}
