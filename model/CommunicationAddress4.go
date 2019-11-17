package model

// Communication device number or electronic address used for communication.
type CommunicationAddress4 struct {

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`

	// Address for the Universal Resource Locator (URL), eg, used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`
}

func (c *CommunicationAddress4) SetEmailAddress(value string) {
	c.EmailAddress = (*Max256Text)(&value)
}

func (c *CommunicationAddress4) SetURLAddress(value string) {
	c.URLAddress = (*Max256Text)(&value)
}
