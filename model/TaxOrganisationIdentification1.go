package model

// Party to which the TaxReport is delivered. This message block contains party details for a specific tax authority.
type TaxOrganisationIdentification1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress6 `xml:"PstlAdr,omitempty"`

	// Set of elements used to indicate how to contact the party.
	ContactDetails *ContactDetails2 `xml:"CtctDtls,omitempty"`
}

func (t *TaxOrganisationIdentification1) SetName(value string) {
	t.Name = (*Max140Text)(&value)
}

func (t *TaxOrganisationIdentification1) AddPostalAddress() *PostalAddress6 {
	t.PostalAddress = new(PostalAddress6)
	return t.PostalAddress
}

func (t *TaxOrganisationIdentification1) AddContactDetails() *ContactDetails2 {
	t.ContactDetails = new(ContactDetails2)
	return t.ContactDetails
}
