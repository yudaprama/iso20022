package model

// Payment terminal or ATM performing the transaction.
type CardAcceptorTerminal1 struct {

	// Identification of the terminal.
	// It correspond to the ISO 8583 field number 41.
	Identification *GenericIdentification32 `xml:"Id"`

	// Location of the terminal.
	Location *PostalAddress18 `xml:"Lctn,omitempty"`

	// Capabilities of the terminal performing the transaction.
	Capabilities *PointOfInteractionCapabilities4 `xml:"Cpblties"`
}

func (c *CardAcceptorTerminal1) AddIdentification() *GenericIdentification32 {
	c.Identification = new(GenericIdentification32)
	return c.Identification
}

func (c *CardAcceptorTerminal1) AddLocation() *PostalAddress18 {
	c.Location = new(PostalAddress18)
	return c.Location
}

func (c *CardAcceptorTerminal1) AddCapabilities() *PointOfInteractionCapabilities4 {
	c.Capabilities = new(PointOfInteractionCapabilities4)
	return c.Capabilities
}
