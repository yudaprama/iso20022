package model

// Choice of formats for a type of price.
type TypeOfPrice46Choice struct {

	// Type expressed as a code.
	Code *TypeOfPrice10Code `xml:"Cd"`

	// Type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TypeOfPrice46Choice) SetCode(value string) {
	t.Code = (*TypeOfPrice10Code)(&value)
}

func (t *TypeOfPrice46Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
