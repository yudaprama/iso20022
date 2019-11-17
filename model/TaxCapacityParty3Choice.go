package model

// Choice of format for the tax capacity information.
type TaxCapacityParty3Choice struct {

	// Party tax capacity expressed as an ISO 20022 code.
	Code *TaxLiability1Code `xml:"Cd"`

	// Party tax capacity expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (t *TaxCapacityParty3Choice) SetCode(value string) {
	t.Code = (*TaxLiability1Code)(&value)
}

func (t *TaxCapacityParty3Choice) AddProprietary() *GenericIdentification38 {
	t.Proprietary = new(GenericIdentification38)
	return t.Proprietary
}
