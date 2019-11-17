package model

// Identification of an entity.
type GenericIdentification76 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3)
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification76) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification76) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification76) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification76) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification76) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
