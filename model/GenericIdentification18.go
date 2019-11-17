package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification18 struct {

	// Identification assigned by an institution.
	Identification *RestrictedFINXMax30Text `xml:"Id"`

	// Name of the identification scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr,omitempty"`
}

func (g *GenericIdentification18) SetIdentification(value string) {
	g.Identification = (*RestrictedFINXMax30Text)(&value)
}

func (g *GenericIdentification18) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification18) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}
