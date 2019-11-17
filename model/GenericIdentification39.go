package model

// Identification expressed as a proprietary type and narrative description.
type GenericIdentification39 struct {

	// Proprietary information issued by the data source scheme issuer.
	Identification *RestrictedFINMax30Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *RestrictedFINMax8Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification39) SetIdentification(value string) {
	g.Identification = (*RestrictedFINMax30Text)(&value)
}

func (g *GenericIdentification39) SetIssuer(value string) {
	g.Issuer = (*RestrictedFINMax8Text)(&value)
}
