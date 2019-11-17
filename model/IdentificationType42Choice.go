package model

// Choice between a code and a data source scheme to specify the type of alternate identification.
type IdentificationType42Choice struct {

	// Type of identification is defined using a code.
	Code *TypeOfIdentification1Code `xml:"Cd"`

	// Type of identification is defined using a data source scheme.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (i *IdentificationType42Choice) SetCode(value string) {
	i.Code = (*TypeOfIdentification1Code)(&value)
}

func (i *IdentificationType42Choice) AddProprietary() *GenericIdentification30 {
	i.Proprietary = new(GenericIdentification30)
	return i.Proprietary
}
