package model

// Communication device number or electronic address used for communication.
type CommunicationAddress6 struct {

	// Type of communication address.
	AddressType *AddressType1Choice `xml:"AdrTp,omitempty"`

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

	// Address for the Universal Resource Locator (URL), for example, used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`
}

func (c *CommunicationAddress6) AddAddressType() *AddressType1Choice {
	c.AddressType = new(AddressType1Choice)
	return c.AddressType
}

func (c *CommunicationAddress6) SetEmail(value string) {
	c.Email = (*Max256Text)(&value)
}

func (c *CommunicationAddress6) SetPhone(value string) {
	c.Phone = (*PhoneNumber)(&value)
}

func (c *CommunicationAddress6) SetMobile(value string) {
	c.Mobile = (*PhoneNumber)(&value)
}

func (c *CommunicationAddress6) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *CommunicationAddress6) SetTelexAddress(value string) {
	c.TelexAddress = (*Max35Text)(&value)
}

func (c *CommunicationAddress6) SetURLAddress(value string) {
	c.URLAddress = (*Max256Text)(&value)
}
