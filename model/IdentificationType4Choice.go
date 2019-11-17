package model

// Choice between a code and a data source scheme to specify the type of alternate identification.
type IdentificationType4Choice struct {

	// Type of identification is defined using a code.
	Code *TypeOfIdentification1Code `xml:"Cd"`

	// Type of identification is defined using a data source scheme.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (i *IdentificationType4Choice) SetCode(value string) {
	i.Code = (*TypeOfIdentification1Code)(&value)
}

func (i *IdentificationType4Choice) AddProprietary() *GenericIdentification20 {
	i.Proprietary = new(GenericIdentification20)
	return i.Proprietary
}
