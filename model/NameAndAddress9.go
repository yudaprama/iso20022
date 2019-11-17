package model

// Information that locates and identifies a party.
type NameAndAddress9 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of a party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`
}

func (n *NameAndAddress9) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress9) AddAddress() *LongPostalAddress2Choice {
	n.Address = new(LongPostalAddress2Choice)
	return n.Address
}
