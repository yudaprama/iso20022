package model

// Identification of an entity.
type GenericIdentification92 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType5Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType6Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Access information to reach the target host.
	RemoteAccess *NetworkParameters5 `xml:"RmotAccs,omitempty"`
}

func (g *GenericIdentification92) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification92) SetType(value string) {
	g.Type = (*PartyType5Code)(&value)
}

func (g *GenericIdentification92) SetIssuer(value string) {
	g.Issuer = (*PartyType6Code)(&value)
}

func (g *GenericIdentification92) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification92) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

func (g *GenericIdentification92) AddRemoteAccess() *NetworkParameters5 {
	g.RemoteAccess = new(NetworkParameters5)
	return g.RemoteAccess
}
