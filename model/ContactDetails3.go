package model

// Communication device number or electronic address used for communication.
type ContactDetails3 struct {

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max140Text `xml:"Nm,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb,omitempty"`

	// Collection of information that identifies a mobile phone number, as defined by telecom services.
	MobileNumber *PhoneNumber `xml:"MobNb,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max2048Text `xml:"EmailAdr,omitempty"`

	// Contact details in an other form.
	Other []*OtherContact1 `xml:"Othr,omitempty"`

	// Preferred method used to reach the technical contact.
	PreferredMethod *PreferredContactMethod1Code `xml:"PrefrdMtd,omitempty"`
}

func (c *ContactDetails3) SetNamePrefix(value string) {
	c.NamePrefix = (*NamePrefix1Code)(&value)
}

func (c *ContactDetails3) SetName(value string) {
	c.Name = (*Max140Text)(&value)
}

func (c *ContactDetails3) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *ContactDetails3) SetMobileNumber(value string) {
	c.MobileNumber = (*PhoneNumber)(&value)
}

func (c *ContactDetails3) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *ContactDetails3) SetEmailAddress(value string) {
	c.EmailAddress = (*Max2048Text)(&value)
}

func (c *ContactDetails3) AddOther() *OtherContact1 {
	newValue := new(OtherContact1)
	c.Other = append(c.Other, newValue)
	return newValue
}

func (c *ContactDetails3) SetPreferredMethod(value string) {
	c.PreferredMethod = (*PreferredContactMethod1Code)(&value)
}
