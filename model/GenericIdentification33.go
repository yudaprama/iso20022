package model

// Identification of an entity.
type GenericIdentification33 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification33) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification33) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification33) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification33) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
