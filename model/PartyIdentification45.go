package model

// Set of elements used to identify a person or an organisation.
type PartyIdentification45 struct {

	// Unique identification, as assigned by an organisation, to unambiguously identify a party.
	Identification *Party8Choice `xml:"Id,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Contact defined for this party.
	ContactDetails []*Contacts3 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification45) AddIdentification() *Party8Choice {
	p.Identification = new(Party8Choice)
	return p.Identification
}

func (p *PartyIdentification45) SetName(value string) {
	p.Name = (*Max35Text)(&value)
}

func (p *PartyIdentification45) AddPostalAddress() *PostalAddress6 {
	p.PostalAddress = new(PostalAddress6)
	return p.PostalAddress
}

func (p *PartyIdentification45) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification45) AddContactDetails() *Contacts3 {
	newValue := new(Contacts3)
	p.ContactDetails = append(p.ContactDetails, newValue)
	return newValue
}
