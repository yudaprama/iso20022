package model

// Choice between a code and a data source scheme to specify the type of alternate identification.
type IdentificationType43Choice struct {

	// Type of identification is defined using a code.
	Code *TypeOfIdentification2Code `xml:"Cd"`

	// Type of identification is defined using a data source scheme.
	Proprietary *GenericIdentification36 `xml:"Prtry"`
}

func (i *IdentificationType43Choice) SetCode(value string) {
	i.Code = (*TypeOfIdentification2Code)(&value)
}

func (i *IdentificationType43Choice) AddProprietary() *GenericIdentification36 {
	i.Proprietary = new(GenericIdentification36)
	return i.Proprietary
}
