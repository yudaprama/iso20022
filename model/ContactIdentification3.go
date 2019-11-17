package model

// Information needed to contact a physical person.
type ContactIdentification3 struct {

	// Business Identifier Code to identify the financial institution that the contact person belongs to.
	BIC *BICIdentifier `xml:"BIC"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm"`

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Role of the party in the activity.
	Role *Max35Text `xml:"Role,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`
}

func (c *ContactIdentification3) SetBIC(value string) {
	c.BIC = (*BICIdentifier)(&value)
}

func (c *ContactIdentification3) SetName(value string) {
	c.Name = (*Max35Text)(&value)
}

func (c *ContactIdentification3) SetNamePrefix(value string) {
	c.NamePrefix = (*NamePrefix1Code)(&value)
}

func (c *ContactIdentification3) SetGivenName(value string) {
	c.GivenName = (*Max35Text)(&value)
}

func (c *ContactIdentification3) SetRole(value string) {
	c.Role = (*Max35Text)(&value)
}

func (c *ContactIdentification3) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification3) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *ContactIdentification3) SetEmailAddress(value string) {
	c.EmailAddress = (*Max256Text)(&value)
}
