package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification2 struct {

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification2) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification2) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
