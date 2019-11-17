package model

// Identification of a person, a financial institution or a non-financial institution.
type PartyIdentification8 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous way of identifying an organisation or an individual person.
	Identification *Party2Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`
}

func (p *PartyIdentification8) SetName(value string) {
	p.Name = (*Max70Text)(&value)
}

func (p *PartyIdentification8) AddPostalAddress() *PostalAddress1 {
	p.PostalAddress = new(PostalAddress1)
	return p.PostalAddress
}

func (p *PartyIdentification8) AddIdentification() *Party2Choice {
	p.Identification = new(Party2Choice)
	return p.Identification
}

func (p *PartyIdentification8) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}
