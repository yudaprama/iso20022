package model

// Identification of a person or an organisation.
type PartyIdentification77 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress19 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *Party11Choice `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification77) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification77) AddPostalAddress() *PostalAddress19 {
	p.PostalAddress = new(PostalAddress19)
	return p.PostalAddress
}

func (p *PartyIdentification77) AddIdentification() *Party11Choice {
	p.Identification = new(Party11Choice)
	return p.Identification
}

func (p *PartyIdentification77) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification77) AddContactDetails() *ContactDetails2 {
	p.ContactDetails = new(ContactDetails2)
	return p.ContactDetails
}
