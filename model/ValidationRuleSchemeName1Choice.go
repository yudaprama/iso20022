package model

// Identifies a name of the identification scheme.
type ValidationRuleSchemeName1Choice struct {

	// Name of the identification scheme, in a coded form as published in an external list.
	Code *ExternalValidationRuleIdentification1Code `xml:"Cd"`

	// Name of the identification scheme, in a free text form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (v *ValidationRuleSchemeName1Choice) SetCode(value string) {
	v.Code = (*ExternalValidationRuleIdentification1Code)(&value)
}

func (v *ValidationRuleSchemeName1Choice) SetProprietary(value string) {
	v.Proprietary = (*Max35Text)(&value)
}
