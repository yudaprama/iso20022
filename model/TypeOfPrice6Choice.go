package model

// Choice of format for the type of price.
type TypeOfPrice6Choice struct {

	// Type of price expressed as an ISO 20022 code.
	Code *TypeOfPrice16Code `xml:"Cd"`

	// Type of price expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (t *TypeOfPrice6Choice) SetCode(value string) {
	t.Code = (*TypeOfPrice16Code)(&value)
}

func (t *TypeOfPrice6Choice) AddProprietary() *GenericIdentification20 {
	t.Proprietary = new(GenericIdentification20)
	return t.Proprietary
}
