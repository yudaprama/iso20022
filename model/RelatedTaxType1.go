package model

// Specifies the type of related tax.
type RelatedTaxType1 struct {

	// Specifies the type of tax.
	TaxType *TaxType3FormatChoice `xml:"TaxTp"`

	// The value of the related tax expressed as an amount.
	Amount *ActiveCurrencyAndAmount `xml:"Amt"`
}

func (r *RelatedTaxType1) AddTaxType() *TaxType3FormatChoice {
	r.TaxType = new(TaxType3FormatChoice)
	return r.TaxType
}

func (r *RelatedTaxType1) SetAmount(value, currency string) {
	r.Amount = NewActiveCurrencyAndAmount(value, currency)
}
