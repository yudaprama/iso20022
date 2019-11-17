package model

// Information that locates and identifies a party.
type NameAndAddress12 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *RestrictedFINXMax140Text `xml:"Nm"`
}

func (n *NameAndAddress12) SetName(value string) {
	n.Name = (*RestrictedFINXMax140Text)(&value)
}
