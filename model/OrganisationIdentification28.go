package model

// Unique identification, as assigned by an organisation, to unambiguously identify a party.
type OrganisationIdentification28 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Unique and unambiguous identification of a party.
	Identification *OrganisationIdentification8 `xml:"Id,omitempty"`

	// Country in which a person resides (the place of a person's home). In the case of a company, it is the country from which the affairs of that company are directed.
	CountryOfResidence *CountryCode `xml:"CtryOfRes,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls"`
}

func (o *OrganisationIdentification28) SetName(value string) {
	o.Name = (*Max140Text)(&value)
}

func (o *OrganisationIdentification28) AddPostalAddress() *PostalAddress6 {
	o.PostalAddress = new(PostalAddress6)
	return o.PostalAddress
}

func (o *OrganisationIdentification28) AddIdentification() *OrganisationIdentification8 {
	o.Identification = new(OrganisationIdentification8)
	return o.Identification
}

func (o *OrganisationIdentification28) SetCountryOfResidence(value string) {
	o.CountryOfResidence = (*CountryCode)(&value)
}

func (o *OrganisationIdentification28) AddContactDetails() *ContactDetails2 {
	o.ContactDetails = new(ContactDetails2)
	return o.ContactDetails
}
