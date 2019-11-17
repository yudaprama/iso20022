package model

// Entity involved in an activity.
type NameAndAddress2 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm"`

	// Information that locates and identifies a specific address, as defined by postal services.
	Address *LongPostalAddress1Choice `xml:"Adr,omitempty"`
}

func (n *NameAndAddress2) SetName(value string) {
	n.Name = (*Max35Text)(&value)
}

func (n *NameAndAddress2) AddAddress() *LongPostalAddress1Choice {
	n.Address = new(LongPostalAddress1Choice)
	return n.Address
}
