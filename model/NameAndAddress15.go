package model

// Information that locates and identifies a party.
type NameAndAddress15 struct {

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of the party.
	PostalAddress *PostalAddress21 `xml:"PstlAdr,omitempty"`
}

func (n *NameAndAddress15) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress15) AddPostalAddress() *PostalAddress21 {
	n.PostalAddress = new(PostalAddress21)
	return n.PostalAddress
}
