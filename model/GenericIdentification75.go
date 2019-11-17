package model

// Identification of an entity.
type GenericIdentification75 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Type of identified entity.
	Type *PartyType11Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType9Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification75) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification75) SetType(value string) {
	g.Type = (*PartyType11Code)(&value)
}

func (g *GenericIdentification75) SetIssuer(value string) {
	g.Issuer = (*PartyType9Code)(&value)
}

func (g *GenericIdentification75) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification75) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
