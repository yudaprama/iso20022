package model

// Communication device number or electronic address used for communication.
type ContactInformation1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	TelephoneNumber *PhoneNumber `xml:"TelNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`
}

func (c *ContactInformation1) SetName(value string) {
	c.Name = (*Max350Text)(&value)
}

func (c *ContactInformation1) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *ContactInformation1) SetTelephoneNumber(value string) {
	c.TelephoneNumber = (*PhoneNumber)(&value)
}

func (c *ContactInformation1) SetEmailAddress(value string) {
	c.EmailAddress = (*Max256Text)(&value)
}
