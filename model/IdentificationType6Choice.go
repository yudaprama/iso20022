package model

// Choice between a code and a data source scheme to specify the type of alternate identification.
type IdentificationType6Choice struct {

	// Type of identification is defined using an ISO 20022 code.
	Code *TypeOfIdentification1Code `xml:"Cd"`

	// Type of identification is defined using a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (i *IdentificationType6Choice) SetCode(value string) {
	i.Code = (*TypeOfIdentification1Code)(&value)
}

func (i *IdentificationType6Choice) AddProprietary() *GenericIdentification30 {
	i.Proprietary = new(GenericIdentification30)
	return i.Proprietary
}
