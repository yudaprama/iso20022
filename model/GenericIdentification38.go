package model

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification38 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification38) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification38) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification38) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}
