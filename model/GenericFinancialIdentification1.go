package model

// Information related to an identification of a financial institution.
type GenericFinancialIdentification1 struct {

	// Unique and unambiguous identification of a person.
	Identification *Max35Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *FinancialIdentificationSchemeName1Choice `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericFinancialIdentification1) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericFinancialIdentification1) AddSchemeName() *FinancialIdentificationSchemeName1Choice {
	g.SchemeName = new(FinancialIdentificationSchemeName1Choice)
	return g.SchemeName
}

func (g *GenericFinancialIdentification1) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
