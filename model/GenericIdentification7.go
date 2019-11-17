package model

// Information expressed in a proprietary manner.
type GenericIdentification7 struct {

	// Entity that assigns the identification.
	Issuer *Max8Text `xml:"Issr"`

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Information *Max35Text `xml:"Inf"`
}

func (g *GenericIdentification7) SetIssuer(value string) {
	g.Issuer = (*Max8Text)(&value)
}

func (g *GenericIdentification7) SetInformation(value string) {
	g.Information = (*Max35Text)(&value)
}
