package model

// Identification of an entity.
type GenericIdentification32 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification32) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification32) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification32) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification32) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
