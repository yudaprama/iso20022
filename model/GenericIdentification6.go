package model

// Balance expressed with a data source scheme.
type GenericIdentification6 struct {

	// Entity that assigns the identification.
	Issuer *Max8Text `xml:"Issr"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Information *Exact4AlphaNumericText `xml:"Inf"`

	// Value of the balance.
	Balance *Number `xml:"Bal"`
}

func (g *GenericIdentification6) SetIssuer(value string) {
	g.Issuer = (*Max8Text)(&value)
}

func (g *GenericIdentification6) SetInformation(value string) {
	g.Information = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification6) SetBalance(value string) {
	g.Balance = (*Number)(&value)
}
