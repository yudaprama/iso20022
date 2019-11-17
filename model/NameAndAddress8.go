package model

// Information that locates and identifies a party.
type NameAndAddress8 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr,omitempty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	AlternativeIdentifier []*Max35Text `xml:"AltrntvIdr,omitempty"`
}

func (n *NameAndAddress8) SetName(value string) {
	n.Name = (*Max350Text)(&value)
}

func (n *NameAndAddress8) AddAddress() *PostalAddress1 {
	n.Address = new(PostalAddress1)
	return n.Address
}

func (n *NameAndAddress8) AddAlternativeIdentifier(value string) {
	n.AlternativeIdentifier = append(n.AlternativeIdentifier, (*Max35Text)(&value))
}
