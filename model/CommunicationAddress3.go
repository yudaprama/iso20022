package model

// Communication device number or electronic address used for communication.
type CommunicationAddress3 struct {

	// Address for electronic mail (e-mail).
	Email *Max256Text `xml:"Email,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	Phone *PhoneNumber `xml:"Phne,omitempty"`

	// Collection of information that identifies a mobile phone number, as defined by telecom services.
	Mobile *PhoneNumber `xml:"Mob,omitempty"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for a telex machine.
	TelexAddress *Max35Text `xml:"TlxAdr,omitempty"`

	// Address for the Universal Resource Locator (URL), eg, used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`
}

func (c *CommunicationAddress3) SetEmail(value string) {
	c.Email = (*Max256Text)(&value)
}

func (c *CommunicationAddress3) SetPhone(value string) {
	c.Phone = (*PhoneNumber)(&value)
}

func (c *CommunicationAddress3) SetMobile(value string) {
	c.Mobile = (*PhoneNumber)(&value)
}

func (c *CommunicationAddress3) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *CommunicationAddress3) SetTelexAddress(value string) {
	c.TelexAddress = (*Max35Text)(&value)
}

func (c *CommunicationAddress3) SetURLAddress(value string) {
	c.URLAddress = (*Max256Text)(&value)
}
