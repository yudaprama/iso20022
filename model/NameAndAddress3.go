package model

// Information that locates and identifies a party.
type NameAndAddress3 struct {

	// Name by which a party is known and is usually used to identify that identity.
	Name *Max70Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr"`
}

func (n *NameAndAddress3) SetName(value string) {
	n.Name = (*Max70Text)(&value)
}

func (n *NameAndAddress3) AddAddress() *PostalAddress1 {
	n.Address = new(PostalAddress1)
	return n.Address
}
