package model

// Choice of formats for the type of tax.
type TaxType1Choice struct {

	// Type of tax expressed as a code.
	Code *TaxType16Code `xml:"Cd"`

	// Type of tax expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxType1Choice) SetCode(value string) {
	t.Code = (*TaxType16Code)(&value)
}

func (t *TaxType1Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
