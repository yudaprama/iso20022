package model

// Choice between a code and a data source scheme to specify the type of alternate identification.
type IdentificationType41Choice struct {

	// Type of identification is defined using a code.
	Code *TypeOfIdentification1Code `xml:"Cd"`

	// Type of identification is defined using a data source scheme.
	Proprietary *GenericIdentification38 `xml:"Prtry"`
}

func (i *IdentificationType41Choice) SetCode(value string) {
	i.Code = (*TypeOfIdentification1Code)(&value)
}

func (i *IdentificationType41Choice) AddProprietary() *GenericIdentification38 {
	i.Proprietary = new(GenericIdentification38)
	return i.Proprietary
}
