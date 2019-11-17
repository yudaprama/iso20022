package model

// Communication device number or electronic address used for communication.
type ContactAttributes1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr"`

	// Collection of information that identifies a phone number, as defined by telecom services.
	PhoneNumber *PhoneNumber `xml:"PhneNb"`

	// Collection of information that identifies a FAX number, as defined by telecom services.
	FaxNumber *PhoneNumber `xml:"FaxNb,omitempty"`

	// Address for electronic mail (e-mail).
	EmailAddress *Max256Text `xml:"EmailAdr,omitempty"`

	// Address for the Universal Resource Locator (URL), eg, used over the www (HTTP) service.
	URLAddress *Max2048Text `xml:"URLAdr,omitempty"`

	// Unique and unambiguous identification of a financial institution, as assigned under a globally recognised or proprietary identification scheme.
	Identification *BICIdentifier `xml:"Id,omitempty"`
}

func (c *ContactAttributes1) SetName(value string) {
	c.Name = (*Max350Text)(&value)
}

func (c *ContactAttributes1) AddPostalAddress() *PostalAddress1 {
	c.PostalAddress = new(PostalAddress1)
	return c.PostalAddress
}

func (c *ContactAttributes1) SetPhoneNumber(value string) {
	c.PhoneNumber = (*PhoneNumber)(&value)
}

func (c *ContactAttributes1) SetFaxNumber(value string) {
	c.FaxNumber = (*PhoneNumber)(&value)
}

func (c *ContactAttributes1) SetEmailAddress(value string) {
	c.EmailAddress = (*Max256Text)(&value)
}

func (c *ContactAttributes1) SetURLAddress(value string) {
	c.URLAddress = (*Max2048Text)(&value)
}

func (c *ContactAttributes1) SetIdentification(value string) {
	c.Identification = (*BICIdentifier)(&value)
}
