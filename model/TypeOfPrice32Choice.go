package model

// Choice of format for the type of price.
type TypeOfPrice32Choice struct {

	// Type of price expressed as an ISO 20022 code.
	Code *TypeOfPrice14Code `xml:"Cd"`

	// Type of price expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TypeOfPrice32Choice) SetCode(value string) {
	t.Code = (*TypeOfPrice14Code)(&value)
}

func (t *TypeOfPrice32Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
