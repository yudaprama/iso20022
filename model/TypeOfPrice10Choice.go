package model

// Choice of format for the type of price.
type TypeOfPrice10Choice struct {

	// Type of price expressed as an ISO 20022 code.
	Code *TypeOfPrice3Code `xml:"Cd"`

	// Type of price expressed as a proprietary code.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (t *TypeOfPrice10Choice) SetCode(value string) {
	t.Code = (*TypeOfPrice3Code)(&value)
}

func (t *TypeOfPrice10Choice) AddProprietary() *GenericIdentification38 {
	t.Proprietary = new(GenericIdentification38)
	return t.Proprietary
}
