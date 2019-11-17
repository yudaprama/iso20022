package model

// Identification of an entity.
type GenericIdentification72 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType6Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification72) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification72) SetIssuer(value string) {
	g.Issuer = (*PartyType6Code)(&value)
}

func (g *GenericIdentification72) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification72) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
