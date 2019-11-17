package model

// Proprietary information related to a balance.
type GenericIdentification56 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Exact4AlphaNumericText `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr"`

	// Short textual description of the scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`

	// Value of the balance.
	Balance *DecimalNumber `xml:"Bal"`
}

func (g *GenericIdentification56) SetIdentification(value string) {
	g.Identification = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification56) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification56) SetSchemeName(value string) {
	g.SchemeName = (*Max35Text)(&value)
}

func (g *GenericIdentification56) SetBalance(value string) {
	g.Balance = (*DecimalNumber)(&value)
}
