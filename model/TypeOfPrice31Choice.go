package model

// Choice of format for the type of price.
type TypeOfPrice31Choice struct {

	// Type of price expressed as a code.
	Code *TypeOfPrice12Code `xml:"Cd"`

	// Type of price expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TypeOfPrice31Choice) SetCode(value string) {
	t.Code = (*TypeOfPrice12Code)(&value)
}

func (t *TypeOfPrice31Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
