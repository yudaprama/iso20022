package model

// Choice of format for the preference to income.
type PreferenceToIncome2Choice struct {

	// Preference to income expressed as an ISO 20022 code.
	Code *PreferenceToIncome1Code `xml:"Cd"`

	// Preference to income expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (p *PreferenceToIncome2Choice) SetCode(value string) {
	p.Code = (*PreferenceToIncome1Code)(&value)
}

func (p *PreferenceToIncome2Choice) AddProprietary() *GenericIdentification20 {
	p.Proprietary = new(GenericIdentification20)
	return p.Proprietary
}
