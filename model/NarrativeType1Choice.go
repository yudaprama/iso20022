package model

// Choice of format for the type of narrative.
type NarrativeType1Choice struct {

	// Type of narrative.
	//
	Code *ExternalNarrativeType1Code `xml:"Cd"`

	// Type of narrative expressed as a proprietary code.
	Proprietary *GenericIdentification1 `xml:"Prtry"`
}

func (n *NarrativeType1Choice) SetCode(value string) {
	n.Code = (*ExternalNarrativeType1Code)(&value)
}

func (n *NarrativeType1Choice) AddProprietary() *GenericIdentification1 {
	n.Proprietary = new(GenericIdentification1)
	return n.Proprietary
}
