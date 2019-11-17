package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification13 struct {

	// Identification assigned by an institution.
	Identification *Max4AlphaNumericText `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`
}

func (g *GenericIdentification13) SetIdentification(value string) {
	g.Identification = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification13) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification13) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
