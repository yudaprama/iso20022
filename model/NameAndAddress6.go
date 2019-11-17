package model

// Name and address of an institution.
type NameAndAddress6 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Information that locates and identifies a specific address, as defined by postal services.
	Address *PostalAddress2 `xml:"Adr"`
}

func (n *NameAndAddress6) SetName(value string) {
	n.Name = (*Max70Text)(&value)
}

func (n *NameAndAddress6) AddAddress() *PostalAddress2 {
	n.Address = new(PostalAddress2)
	return n.Address
}
