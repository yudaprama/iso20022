package model

// Describes the details of an organisation.
type PartyIdentification58 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Name by which a party is known and which is usually used to identify that party.
	LegalName *Max140Text `xml:"LglNm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress11 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous way to identify the party.
	Identification *Party13Choice `xml:"Id"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Indicate how to contact the party.
	ContactDetails *ContactDetails3 `xml:"CtctDtls,omitempty"`
}

func (p *PartyIdentification58) SetName(value string) {
	p.Name = (*Max140Text)(&value)
}

func (p *PartyIdentification58) SetLegalName(value string) {
	p.LegalName = (*Max140Text)(&value)
}

func (p *PartyIdentification58) AddPostalAddress() *PostalAddress11 {
	p.PostalAddress = new(PostalAddress11)
	return p.PostalAddress
}

func (p *PartyIdentification58) AddIdentification() *Party13Choice {
	p.Identification = new(Party13Choice)
	return p.Identification
}

func (p *PartyIdentification58) SetCountryOfResidence(value string) {
	p.CountryOfResidence = (*CountryCode)(&value)
}

func (p *PartyIdentification58) AddContactDetails() *ContactDetails3 {
	p.ContactDetails = new(ContactDetails3)
	return p.ContactDetails
}
