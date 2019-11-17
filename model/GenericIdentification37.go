package model

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification37 struct {

	// Proprietary information, often a code, issued by the data source scheme issuer.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification37) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification37) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
