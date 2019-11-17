package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification1 struct {

	// Identification assigned by an institution.
	Identification *Max35Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification1) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification1) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification1) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
