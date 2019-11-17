package model

// Identification of an entity.
type GenericIdentification31 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification31) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification31) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification31) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification31) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
