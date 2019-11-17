package model

// Information that locates and identifies a party.
type NameAndAddress10 struct {

	// Name by which a party is known and is usually used to identify that identity.
	Name *Max140Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress6 `xml:"Adr"`
}

func (n *NameAndAddress10) SetName(value string) {
	n.Name = (*Max140Text)(&value)
}

func (n *NameAndAddress10) AddAddress() *PostalAddress6 {
	n.Address = new(PostalAddress6)
	return n.Address
}
