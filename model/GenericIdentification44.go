package model

// Information related to the identification of an individual person.
type GenericIdentification44 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	Type *OtherIdentification1Choice `xml:"Tp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`
}

func (g *GenericIdentification44) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification44) AddType() *OtherIdentification1Choice {
	g.Type = new(OtherIdentification1Choice)
	return g.Type
}

func (g *GenericIdentification44) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification44) SetIssueDate(value string) {
	g.IssueDate = (*ISODate)(&value)
}

func (g *GenericIdentification44) SetExpiryDate(value string) {
	g.ExpiryDate = (*ISODate)(&value)
}
