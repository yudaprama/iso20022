package model

// Choice between a code and a data source scheme to specify the type of alternate identification.
type IdentificationType40Choice struct {

	// Type of identification is defined using a code.
	Code *TypeOfIdentification2Code `xml:"Cd"`

	// Type of identification is defined using a data source scheme.
	Proprietary *GenericIdentification29 `xml:"Prtry"`
}

func (i *IdentificationType40Choice) SetCode(value string) {
	i.Code = (*TypeOfIdentification2Code)(&value)
}

func (i *IdentificationType40Choice) AddProprietary() *GenericIdentification29 {
	i.Proprietary = new(GenericIdentification29)
	return i.Proprietary
}
