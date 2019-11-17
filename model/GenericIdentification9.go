package model

// Information related to the identification of an individual person.
type GenericIdentification9 struct {

	// Name or number assigned by an entity to enable recognition of that entity, eg, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	IdentificationType *PersonIdentificationType1Code `xml:"IdTp"`

	// Specifies the nature of the identification.
	ExtendedIdentificationType *Extended350Code `xml:"XtndedIdTp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`
}

func (g *GenericIdentification9) SetIdentification(value string) {
	g.Identification = (*Max35Text)(&value)
}

func (g *GenericIdentification9) SetIdentificationType(value string) {
	g.IdentificationType = (*PersonIdentificationType1Code)(&value)
}

func (g *GenericIdentification9) SetExtendedIdentificationType(value string) {
	g.ExtendedIdentificationType = (*Extended350Code)(&value)
}

func (g *GenericIdentification9) SetIssuer(value string) {
	g.Issuer = (*Max35Text)(&value)
}

func (g *GenericIdentification9) SetIssueDate(value string) {
	g.IssueDate = (*ISODate)(&value)
}

func (g *GenericIdentification9) SetExpiryDate(value string) {
	g.ExpiryDate = (*ISODate)(&value)
}
