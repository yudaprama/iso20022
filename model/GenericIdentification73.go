package model

// Identification of an entity.
type GenericIdentification73 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType9Code `xml:"Tp,omitempty"`

	// Entity assigning the identification , for example merchant, acceptor, acquirer or card issuer.
	Issuer *PartyType9Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`
}

func (g *GenericIdentification73) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification73) SetType(value string) {
	g.Type = (*PartyType9Code)(&value)
}

func (g *GenericIdentification73) SetIssuer(value string) {
	g.Issuer = (*PartyType9Code)(&value)
}

func (g *GenericIdentification73) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification73) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}
