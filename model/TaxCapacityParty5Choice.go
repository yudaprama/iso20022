package model

// Choice of format for the tax capacity information.
type TaxCapacityParty5Choice struct {

	// Party tax capacity expressed as an ISO 20022 code.
	Code *TaxLiability1Code `xml:"Cd"`

	// Party tax capacity expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxCapacityParty5Choice) SetCode(value string) {
	t.Code = (*TaxLiability1Code)(&value)
}

func (t *TaxCapacityParty5Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
