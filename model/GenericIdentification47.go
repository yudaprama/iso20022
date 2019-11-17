package model

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification47 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (g *GenericIdentification47) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification47) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification47) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}
