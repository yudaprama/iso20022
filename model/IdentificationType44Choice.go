package model

// Choice between a code and a data source scheme to specify the type of alternate identification.
type IdentificationType44Choice struct {

	// Type of identification is defined using a code.
	Code *TypeOfIdentification1Code `xml:"Cd"`

	// Type of identification is defined using a data source scheme.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (i *IdentificationType44Choice) SetCode(value string) {
	i.Code = (*TypeOfIdentification1Code)(&value)
}

func (i *IdentificationType44Choice) AddProprietary() *GenericIdentification47 {
	i.Proprietary = new(GenericIdentification47)
	return i.Proprietary
}
