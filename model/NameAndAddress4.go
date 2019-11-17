package model

// Information that locates and identifies a party.
type NameAndAddress4 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr"`
}

func (n *NameAndAddress4) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress4) AddAddress() *PostalAddress1 {
	n.Address = new(PostalAddress1)
	return n.Address
}
