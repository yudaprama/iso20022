package model

// Choice of format for the type of price.
type TypeOfPrice27Choice struct {

	// Type of price expressed as an ISO 20022 code.
	Code *TypeOfPrice30Code `xml:"Cd"`

	// Type of price expressed as a proprietary code.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (t *TypeOfPrice27Choice) SetCode(value string) {
	t.Code = (*TypeOfPrice30Code)(&value)
}

func (t *TypeOfPrice27Choice) AddProprietary() *GenericIdentification29 {
	t.Proprietary = new(GenericIdentification29)
	return t.Proprietary
}
