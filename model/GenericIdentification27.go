package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification27 struct {

	// Identification assigned by an institution.
	Identification *Max4AlphaNumericText `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`
}

func (g *GenericIdentification27) SetIdentification(value string) {
	g.Identification = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification27) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification27) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}
