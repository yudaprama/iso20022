package model

// Information that locates and identifies a specific address, as defined by postal services.
type NameAndAddress1 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Postal address of a party.
	Address *LongPostalAddress1Choice `xml:"Adr"`
}

func (n *NameAndAddress1) SetName(value string) {
	n.Name = (*Max35Text)(&value)
}

func (n *NameAndAddress1) AddAddress() *LongPostalAddress1Choice {
	n.Address = new(LongPostalAddress1Choice)
	return n.Address
}
