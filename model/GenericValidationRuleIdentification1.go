package model

// Information for the identification of a validation rule.
type GenericValidationRuleIdentification1 struct {

	// Unique and unambiguous identification of a validation rule.
	Identification *Max35Text `xml:"Id"`

	// Further information on the validation rule as identified in the Identification.
	Description *Max350Text `xml:"Desc,omitempty"`

	// Name of the identification scheme.
	SchemeName *ValidationRuleSchemeName1Choice `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericValidationRuleIdentification1) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericValidationRuleIdentification1) SetDescription(value string) {
	g.Description = (*Max350Text)(&value)
}

func (g *GenericValidationRuleIdentification1) AddSchemeName() *ValidationRuleSchemeName1Choice {
	g.SchemeName = new(ValidationRuleSchemeName1Choice)
	return g.SchemeName
}

func (g *GenericValidationRuleIdentification1) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
