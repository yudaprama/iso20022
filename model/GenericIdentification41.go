package model

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification41 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification41) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification41) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification41) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}
