package model

// Identification of an entity.
type GenericIdentification35 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType5Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType6Code `xml:"Issr,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification35) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification35) SetType(value string) {
	g.Type = (*PartyType5Code)(&value)
}

func (g *GenericIdentification35) SetIssuer(value string) {
	g.Issuer = (*PartyType6Code)(&value)
}

func (g *GenericIdentification35) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
