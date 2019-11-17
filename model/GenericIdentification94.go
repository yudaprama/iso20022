package model

// Identification of an entity.
type GenericIdentification94 struct {

	// Identification of the entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identified entity.
	Type *PartyType3Code `xml:"Tp,omitempty"`

	// Entity assigning the identification  (for example merchant, acceptor, acquirer, or tax authority).
	Issuer *PartyType4Code `xml:"Issr,omitempty"`

	// Country of the entity (ISO 3166-1 alpha-2 or alpha-3).
	Country *Min2Max3AlphaText `xml:"Ctry,omitempty"`

	// Name of the entity.
	ShortName *Max35Text `xml:"ShrtNm,omitempty"`

	// Access information to reach the target host.
	RemoteAccess *NetworkParameters5 `xml:"RmotAccs,omitempty"`
}

func (g *GenericIdentification94) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification94) SetType(value string) {
	g.Type = (*PartyType3Code)(&value)
}

func (g *GenericIdentification94) SetIssuer(value string) {
	g.Issuer = (*PartyType4Code)(&value)
}

func (g *GenericIdentification94) SetCountry(value string) {
	g.Country = (*Min2Max3AlphaText)(&value)
}

func (g *GenericIdentification94) SetShortName(value string) {
	g.ShortName = (*Max35Text)(&value)
}

func (g *GenericIdentification94) AddRemoteAccess() *NetworkParameters5 {
	g.RemoteAccess = new(NetworkParameters5)
	return g.RemoteAccess
}
