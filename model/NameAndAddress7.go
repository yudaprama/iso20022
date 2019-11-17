package model

// Name and address of an institution.
type NameAndAddress7 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max70Text `xml:"Nm"`

	// Information that locates and identifies a specific address, as defined by postal services.
	PostalAddress *PostalAddress1 `xml:"PstlAdr"`
}

func (n *NameAndAddress7) SetName(value string) {
	n.Name = (*Max70Text)(&value)
}

func (n *NameAndAddress7) AddPostalAddress() *PostalAddress1 {
	n.PostalAddress = new(PostalAddress1)
	return n.PostalAddress
}
