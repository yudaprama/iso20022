package model

// Information related to an identification, eg, party identification or account identification.
type GenericIdentification8 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	IdentificationType *IdentificationType1 `xml:"IdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification8) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification8) AddIdentificationType() *IdentificationType1 {
	g.IdentificationType = new(IdentificationType1)
	return g.IdentificationType
}

func (g *GenericIdentification8) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
