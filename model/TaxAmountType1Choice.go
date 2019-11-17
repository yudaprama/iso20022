package model

// Specifies the amount type.
type TaxAmountType1Choice struct {

	// Specifies the amount type, in a coded form.
	Code *ExternalTaxAmountType1Code `xml:"Cd"`

	// Specifies the amount type, in a free-text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (t *TaxAmountType1Choice) SetCode(value string) {
	t.Code = (*ExternalTaxAmountType1Code)(&value)
}

func (t *TaxAmountType1Choice) SetProprietary(value string) {
	t.Proprietary = (*Max35Text)(&value)
}
