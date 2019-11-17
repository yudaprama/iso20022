package model

// Identification using a proprietary scheme.
type GenericIdentification163 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification163) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification163) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification163) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}
