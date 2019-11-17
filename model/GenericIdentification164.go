package model

// Information related to an identification, for example, party identification or account identification.
type GenericIdentification164 struct {

	// Name or number assigned by an entity to enable recognition of that entity, for example, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identifier.
	IdentificationType *OtherIdentification3Choice `xml:"IdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (g *GenericIdentification164) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification164) AddIdentificationType() *OtherIdentification3Choice {
	g.IdentificationType = new(OtherIdentification3Choice)
	return g.IdentificationType
}

func (g *GenericIdentification164) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}
