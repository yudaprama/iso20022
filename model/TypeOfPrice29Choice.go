package model

// Choice of format for the type of price.
type TypeOfPrice29Choice struct {

	// Type of price expressed as an ISO 20022 code.
	Code *TypeOfPrice14Code `xml:"Cd"`

	// Type of price expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TypeOfPrice29Choice) SetCode(value string) {
	t.Code = (*TypeOfPrice14Code)(&value)
}

func (t *TypeOfPrice29Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
