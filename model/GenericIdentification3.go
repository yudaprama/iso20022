package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification3 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification3) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification3) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
