package model

// Device used by the customer to perform the payment.
type CustomerDevice1 struct {

	// Identifier of the component.
	Identification *Max35Text `xml:"Id,omitempty"`

	// Type of component.
	Type *Max35Text `xml:"Tp,omitempty"`

	// Provider of the component.
	Provider *Max35Text `xml:"Prvdr,omitempty"`
}

func (c *CustomerDevice1) SetIdentification(value string) {
	c.Identification = (*Max35Text)(&value)
}

func (c *CustomerDevice1) SetType(value string) {
	c.Type = (*Max35Text)(&value)
}

func (c *CustomerDevice1) SetProvider(value string) {
	c.Provider = (*Max35Text)(&value)
}
