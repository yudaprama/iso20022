package model

// Proprietary information related to a balance.
type GenericIdentification144 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`

	// Value of the balance.
	Balance *RestrictedFINDecimalNumber `xml:"Bal"`
}

func (g *GenericIdentification144) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification144) SetIssuer(value string) {
	g.Issuer = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification144) SetSchemeName(value string) {
	g.SchemeName = (*Max4AlphaNumericText)(&value)
}

func (g *GenericIdentification144) SetBalance(value string) {
	g.Balance = (*RestrictedFINDecimalNumber)(&value)
}
