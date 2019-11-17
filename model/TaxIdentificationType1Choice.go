package model

// Choice of formats for the type of tax identification.
type TaxIdentificationType1Choice struct {

	// Type of tax identification number or identifier expressed as a code.
	Code *TaxIdentificationNumberType1Code `xml:"Cd"`

	// Type of tax identification number expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxIdentificationType1Choice) SetCode(value string) {
	t.Code = (*TaxIdentificationNumberType1Code)(&value)
}

func (t *TaxIdentificationType1Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
