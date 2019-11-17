package model

// Information that locates and identifies a party.
type NameAndAddress13 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress8 `xml:"Adr,omitempty"`
}

func (n *NameAndAddress13) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress13) AddAddress() *PostalAddress8 {
	n.Address = new(PostalAddress8)
	return n.Address
}
