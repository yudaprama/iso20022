package model

// Information related to the identification of a party.
type GenericIdentification82 struct {

	// Name or number assigned by an entity to enable recognition of that entity.
	Identification *Max35Text `xml:"Id"`

	// Type of identification.
	Type *OtherIdentification3Choice `xml:"Tp"`

	// Entity that assigns the identification.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Name of the state, county or country sub-division that issued the identification document.
	State *Max70Text `xml:"Stat,omitempty"`

	// Country that issued the identification document.
	IssuerCountry *CountryCode `xml:"IssrCtry,omitempty"`
}

func (g *GenericIdentification82) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification82) AddType() *OtherIdentification3Choice {
	g.Type = new(OtherIdentification3Choice)
	return g.Type
}

func (g *GenericIdentification82) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification82) SetIssueDate(value string) {
	g.IssueDate = (*ISODate)(&value)
}

func (g *GenericIdentification82) SetExpiryDate(value string) {
	g.ExpiryDate = (*ISODate)(&value)
}

func (g *GenericIdentification82) SetState(value string) {
	g.State = (*Max70Text)(&value)
}

func (g *GenericIdentification82) SetIssuerCountry(value string) {
	g.IssuerCountry = (*CountryCode)(&value)
}
