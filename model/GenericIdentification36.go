package model

// Identification using a proprietary scheme.
type GenericIdentification36 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification36) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification36) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification36) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}
