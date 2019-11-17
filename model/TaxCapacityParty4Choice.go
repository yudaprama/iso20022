package model

// Choice of format for the tax capacity information.
type TaxCapacityParty4Choice struct {

	// Party tax capacity expressed as an ISO 20022 code.
	Code *TaxLiability1Code `xml:"Cd"`

	// Party tax capacity expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TaxCapacityParty4Choice) SetCode(value string) {
	t.Code = (*TaxLiability1Code)(&value)
}

func (t *TaxCapacityParty4Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
