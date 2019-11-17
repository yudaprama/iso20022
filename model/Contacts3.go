package model

// Address for electronic mail (e-mail).
type Contacts3 struct {

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	MobileNumber *PhoneNumber `xml:"MobNb,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max2048Text `xml:"EmailAdr,omitempty"`

	// Contact details in an other form.
	Other *Max35Text `xml:"Othr,omitempty"`

	// Title of the function.
	JobTitle *Max35Text `xml:"JobTitl,omitempty"`

	// Role of a person in an organisation.
	Responsibility *Max35Text `xml:"Rspnsblty,omitempty"`

	// Identification of a division of a large organisation or building.
	Department *Max70Text `xml:"Dept,omitempty"`
}

func (c *Contacts3) SetNamePrefix(value string) {
	c.NamePrefix = (*NamePrefix1Code)(&value)
}

func (c *Contacts3) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *Contacts3) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *Contacts3) SetMobileNumber(value string) {
	c.MobileNumber = (*PhoneNumber)(&value)
}

func (c *Contacts3) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *Contacts3) SetEmailAddress(value string) {
	c.EmailAddress = (*Max2048Text)(&value)
}

func (c *Contacts3) SetOther(value string) {
	c.Other = (*Max35Text)(&value)
}

func (c *Contacts3) SetJobTitle(value string) {
	c.JobTitle = (*Max35Text)(&value)
}

func (c *Contacts3) SetResponsibility(value string) {
	c.Responsibility = (*Max35Text)(&value)
}

func (c *Contacts3) SetDepartment(value string) {
	c.Department = (*Max70Text)(&value)
}
