package model

// Choice of format for the tax capacity information.
type TaxCapacityParty1Choice struct {

	// Party tax capacity expressed as an ISO 20022 code.
	Code *TaxLiability1Code `xml:"Cd"`

	// Party tax capacity expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *TaxCapacityParty1Choice) SetCode(value string) {
	t.Code = (*TaxLiability1Code)(&value)
}

func (t *TaxCapacityParty1Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}
