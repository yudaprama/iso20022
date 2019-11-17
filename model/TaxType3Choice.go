package model

// Choice of formats for the type of tax.
type TaxType3Choice struct {

	// Type of tax expressed as a code.
	Code *TaxType17Code `xml:"Cd"`

	// Type of tax expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxType3Choice) SetCode(value string) {
	t.Code = (*TaxType17Code)(&value)
}

func (t *TaxType3Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
