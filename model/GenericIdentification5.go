package model

// Information expressed in a proprietary manner.
type GenericIdentification5 struct {

	// Entity that assigns the identification.
	Issuer *Max8Text `xml:"Issr"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Information *Exact4AlphaNumericText `xml:"Inf"`

	// Additional information.
	Narrative *Max35Text `xml:"Nrrtv,omitempty"`
}

func (g *GenericIdentification5) SetIssuer(value string) {
	g.Issuer = (*Max8Text)(&value)
}

func (g *GenericIdentification5) SetInformation(value string) {
	g.Information = (*Exact4AlphaNumericText)(&value)
}

func (g *GenericIdentification5) SetNarrative(value string) {
	g.Narrative = (*Max35Text)(&value)
}
