package model

// Choice of format for the linkage type.
type LinkageType3Choice struct {

	// Linkage type expressed as an ISO 20022 code.
	Code *LinkageType1Code `xml:"Cd"`

	// Linkage type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (l *LinkageType3Choice) SetCode(value string) {
	l.Code = (*LinkageType1Code)(&value)
}

func (l *LinkageType3Choice) AddProprietary() *GenericIdentification30 {
	l.Proprietary = new(GenericIdentification30)
	return l.Proprietary
}
