package model

// Information related to a generic account identification.
type GenericAccountIdentification1 struct {

	// Identification assigned by an institution.
	Identification *Max34Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *AccountSchemeName1Choice `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericAccountIdentification1) SetIdentification(value string) {
	g.Identification = (*Max34Text)(&value)
}

func (g *GenericAccountIdentification1) AddSchemeName() *AccountSchemeName1Choice {
	g.SchemeName = new(AccountSchemeName1Choice)
	return g.SchemeName
}

func (g *GenericAccountIdentification1) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
