package model

// Information related to the identification of an individual person.
type AlternateIdentification4 struct {

	// Name or number assigned by an entity to enable recognition of that entity, for example, account identifier.
	Identification *Max35Text `xml:"Id"`

	// Specifies the nature of the identification.
	Type *OtherIdentification4Choice `xml:"Tp"`

	// Entity that assigns the identifier.
	Issuer *Max35Text `xml:"Issr,omitempty"`

	// Date at which the identification was issued.
	IssueDate *ISODate `xml:"IsseDt,omitempty"`

	// Date at which the identification expires.
	ExpiryDate *ISODate `xml:"XpryDt,omitempty"`

	// Country that issued the identification document.
	IssuerCountry *CountryCode `xml:"IssrCtry,omitempty"`
}

func (a *AlternateIdentification4) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AlternateIdentification4) AddType() *OtherIdentification4Choice {
	a.Type = new(OtherIdentification4Choice)
	return a.Type
}

func (a *AlternateIdentification4) SetIssuer(value string) {
	a.Issuer = (*Max35Text)(&value)
}

func (a *AlternateIdentification4) SetIssueDate(value string) {
	a.IssueDate = (*ISODate)(&value)
}

func (a *AlternateIdentification4) SetExpiryDate(value string) {
	a.ExpiryDate = (*ISODate)(&value)
}

func (a *AlternateIdentification4) SetIssuerCountry(value string) {
	a.IssuerCountry = (*CountryCode)(&value)
}
