package model

// Communication information.
type CommunicationAddress5 struct {

	// Postal address of the entity.
	PostalAddress *PostalAddress18 `xml:"PstlAdr,omitempty"`

	// Address for electronic mail (e-mail).
	Email *Max256Text `xml:"Email,omitempty"`

	// Address for the Universal Resource Locator (URL), for example used over the www (HTTP) service.
	URLAddress *Max256Text `xml:"URLAdr,omitempty"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	Phone *PhoneNumber `xml:"Phne,omitempty"`

	// Phone number of the customer service.
	CustomerService *PhoneNumber `xml:"CstmrSvc,omitempty"`

	// Additional information used to facilitate contact with the card acceptor, for instance sales agent name, dispute manager name.
	AdditionalContactInformation *Max256Text `xml:"AddtlCtctInf,omitempty"`
}

func (c *CommunicationAddress5) AddPostalAddress() *PostalAddress18 {
	c.PostalAddress = new(PostalAddress18)
	return c.PostalAddress
}

func (c *CommunicationAddress5) SetEmail(value string) {
	c.Email = (*Max256Text)(&value)
}

func (c *CommunicationAddress5) SetURLAddress(value string) {
	c.URLAddress = (*Max256Text)(&value)
}

func (c *CommunicationAddress5) SetPhone(value string) {
	c.Phone = (*PhoneNumber)(&value)
}

func (c *CommunicationAddress5) SetCustomerService(value string) {
	c.CustomerService = (*PhoneNumber)(&value)
}

func (c *CommunicationAddress5) SetAdditionalContactInformation(value string) {
	c.AdditionalContactInformation = (*Max256Text)(&value)
}
