package model

// Information related to the identification of an individual person.
type GenericIdentification55 struct {

	// Name or number assigned by an entity to enable recognition of that entity, for example, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	Type *OtherIdentification2Choice `xml:"Tp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Country that issued the identification document.
	IssuerCountry *CountryCode `xml:"IssrCtry,omitempty"`
}

func (g *GenericIdentification55) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification55) AddType() *OtherIdentification2Choice {
	g.Type = new(OtherIdentification2Choice)
	return g.Type
}

func (g *GenericIdentification55) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification55) SetIssueDate(value string) {
	g.IssueDate = (*ISODate)(&value)
}

func (g *GenericIdentification55) SetExpiryDate(value string) {
	g.ExpiryDate = (*ISODate)(&value)
}

func (g *GenericIdentification55) SetIssuerCountry(value string) {
	g.IssuerCountry = (*CountryCode)(&value)
}
