package model

// Choice of format for the type of price.
type TypeOfPrice45Choice struct {

	// Type of price expressed as an ISO 20022 code.
	Code *TypeOfPrice16Code `xml:"Cd"`

	// Type of price expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TypeOfPrice45Choice) SetCode(value string) {
	t.Code = (*TypeOfPrice16Code)(&value)
}

func (t *TypeOfPrice45Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
