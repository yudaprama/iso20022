package model

// Choice of formats for the specification of the name prefix.
type NamePrefix1Choice struct {

	// Name prefix expressed as a code.
	Code *NamePrefix1Code `xml:"Cd"`

	// Name prefix expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (n *NamePrefix1Choice) SetCode(value string) {
	n.Code = (*NamePrefix1Code)(&value)
}

func (n *NamePrefix1Choice) AddProprietary() *GenericIdentification47 {
	n.Proprietary = new(GenericIdentification47)
	return n.Proprietary
}
